package config

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/owu/number-sender/internal/pkg/logger"
	"github.com/owu/number-sender/internal/pkg/model"
	"go.uber.org/zap"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type LoadConfigs struct {
	configs *model.Configs
}

func NewLoadConfigs() *LoadConfigs {
	instance := &LoadConfigs{}
	configs := &model.Configs{}

	// 优先从环境变量获取配置文件路径
	var cfgVar string
	if envCfg := os.Getenv("CONFIG_PATH"); envCfg != "" {
		cfgVar = envCfg
	} else {
		flag.StringVar(&cfgVar, "config", "config/config-dev.toml", "config file path")
		flag.Parse()
	}
	
	cfgFile := instance.verifyFile(cfgVar)
	if _, err := toml.DecodeFile(cfgFile, &configs); err != nil {
		logger.Log.Error("config toml.DecodeFile failed", zap.Error(err), zap.String("cfgFile", cfgFile))
		log.Fatal(err)
	}

	// 允许通过环境变量覆盖端口配置
	if port := os.Getenv("HTTP_PORT"); port != "" {
		if p, err := strconv.Atoi(port); err == nil {
			configs.Server.HTTP.Port = p
			logger.Log.Info("Override HTTP port from environment variable", zap.Int("port", p))
		}
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
		logger.Log.Error("config Getwd failed", zap.Error(err))
		panic("config Getwd failed: " + err.Error())
	}

	if cfgFile = instance.fileInfo(root, cfgVar); len(cfgFile) != 0 {
		return cfgFile
	}

	execute, err = os.Executable()
	if err != nil {
		logger.Log.Error("config Executable failed", zap.Error(err))
		panic("config Executable failed: " + err.Error())
	}

	if cfgFile = instance.fileInfo(filepath.Dir(execute), cfgVar); len(cfgFile) != 0 {
		return cfgFile
	}

	logger.Log.Error("config Get Root failed", zap.String("cfgVar", cfgVar))
	panic("config Get Root failed")
}

func (instance *LoadConfigs) fileInfo(root, cfgVar string) string {
	fullPath := filepath.Join(root, "/", cfgVar)
	if _, err := os.Stat(fullPath); err == nil {
		return fullPath
	}
	return ""
}
