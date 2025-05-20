package config

import (
	"fmt"
	"github.com/owu/number-sender/internal/pkg/dto"
)

func (instance *LoadConfigs) Env() string {
	return instance.configs.App.Env
}

func (instance *LoadConfigs) HttpPort() string {
	return fmt.Sprintf(":%d", instance.configs.Server.HTTP.Port)
}

func (instance *LoadConfigs) RedisDefault() dto.Default {
	return instance.configs.Server.Redis.Default
}

func (instance *LoadConfigs) Encrypt() string {
	return instance.configs.Api.Encrypt
}

func (instance *LoadConfigs) ApiRules() dto.Rules {
	return instance.configs.Api.Rules
}
