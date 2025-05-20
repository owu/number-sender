package config

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/owu/number-sender/internal/pkg/dto"
	"log"
	"os"
	"path/filepath"
)

type LoadConfigs struct {
	configs *dto.Configs
}

func NewLoadConfigs() *LoadConfigs {
	instance := &LoadConfigs{}
	configs := &dto.Configs{}

	var cfgVar string
	flag.StringVar(&cfgVar, "config", "config/config-dev.toml", "config file path")
	flag.Parse()
	cfgFile := instance.verifyFile(cfgVar)
	if _, err := toml.DecodeFile(cfgFile, &configs); err != nil {
		log.Fatal(err)
		panic("config toml.DecodeFile failed," + err.Error())
	}

	instance.configs = configs
	return instance
}

func (instance *LoadConfigs) verifyFile(cfgVar string) string {
	var (
		root, cfgFile, execute string
		err                    error
	)

	if root, err = os.Getwd(); err != nil {
		panic("config Getwd failed")
	}

	if cfgFile = instance.fileInfo(root, cfgVar); len(cfgFile) != 0 {
		return cfgFile
	}

	execute, err = os.Executable()
	if err != nil {
		panic("config Executable failed")
	}

	if cfgFile = instance.fileInfo(filepath.Dir(execute), cfgVar); len(cfgFile) != 0 {
		return cfgFile
	}

	panic("config Get Root failed")
}

func (instance *LoadConfigs) fileInfo(root, cfgVar string) string {
	fullPath := filepath.Join(root, "/", cfgVar)
	if _, err := os.Stat(fullPath); err == nil {
		return fullPath
	}
	return ""
}
