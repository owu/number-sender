package redis

import (
	"context"
	"fmt"
	redisv9 "github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"math/rand"
	"number-sender/internal/pkg/config"
	"number-sender/internal/pkg/logger"
	"time"
)

type DefaultRedis struct {
	instanceMaster *redisv9.Client
	instanceSlaves []*redisv9.Client
}

func NewDefaultRedis(configs *config.LoadConfigs) *DefaultRedis {
	configRedisDefault := configs.RedisDefault()

	dialTimeout := time.Duration(configRedisDefault.DialTimeout) * time.Second
	connMaxIdleTime := time.Duration(configRedisDefault.IdleTimeout) * time.Second
	maxActive := configRedisDefault.MaxActive
	maxIdle := configRedisDefault.MaxIdle
	readTimeout := time.Duration(configRedisDefault.ReadTimeout) * time.Second

	master, err := redisv9.ParseURL(configRedisDefault.Master.Addr)
	if err != nil {
		logger.Log.Error("redis Unable to ParseURL to RedisDefaultMaster", zap.Error(err))
		panic("redis Unable to ParseURL to RedisDefaultMaster: " + err.Error())
	}

	master.DialTimeout = dialTimeout
	master.ConnMaxIdleTime = connMaxIdleTime
	master.MaxActiveConns = maxActive
	master.MaxIdleConns = maxIdle
	master.ReadTimeout = readTimeout
	instanceMaster := redisv9.NewClient(master)
	if err = instanceMaster.Ping(context.Background()).Err(); err != nil {
		logger.Log.Error("redis Unable to connect to RedisDefaultMaster", zap.Error(err))
		panic("redis Unable to connect to RedisDefaultMaster: " + err.Error())
	}

	if len(configRedisDefault.Slaves.Addr) == 0 {
		logger.Log.Error("redis RedisDefaultSlaves is empty")
		panic("redis RedisDefaultSlaves is empty")
	}

	instanceSlaves := make([]*redisv9.Client, 0)

	for _, addr := range configRedisDefault.Slaves.Addr {
		slave, err1 := redisv9.ParseURL(addr)
		if err1 != nil {
			logger.Log.Error("redis Unable to ParseURL to RedisDefaultSlaves", zap.Error(err1), zap.String("addr", addr))
			panic(fmt.Sprintf("redis Unable to ParseURL to RedisDefaultSlaves %s: %v", addr, err1))
		}
		slave.DialTimeout = dialTimeout
		slave.ConnMaxIdleTime = connMaxIdleTime
		slave.MaxActiveConns = maxActive
		slave.MaxIdleConns = maxIdle
		slave.ReadTimeout = readTimeout
		instanceSlave := redisv9.NewClient(slave)
		if err = instanceSlave.Ping(context.Background()).Err(); err != nil {
			logger.Log.Error("redis Unable to connect to RedisDefaultSlaves", zap.Error(err), zap.String("addr", addr))
			panic(fmt.Sprintf("redis Unable to connect to RedisDefaultSlaves %s: %v", addr, err))
		}
		instanceSlaves = append(instanceSlaves, instanceSlave)
	}

	return &DefaultRedis{
		instanceMaster: instanceMaster,
		instanceSlaves: instanceSlaves,
	}
}

func (instance *DefaultRedis) slave() *redisv9.Client {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := rand.Intn(len(instance.instanceSlaves))
	return instance.instanceSlaves[randomIndex]
}

func (instance *DefaultRedis) master() *redisv9.Client {
	return instance.instanceMaster
}
