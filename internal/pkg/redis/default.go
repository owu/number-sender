package redis

import (
	"context"
	"github.com/owu/number-sender/internal/pkg/config"
	redisv9 "github.com/redis/go-redis/v9"
	"math/rand"
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
		panic("redis Unable to ParseURL to RedisDefaultMaster," + err.Error())
	}

	master.DialTimeout = dialTimeout
	master.ConnMaxIdleTime = connMaxIdleTime
	master.MaxActiveConns = maxActive
	master.MaxIdleConns = maxIdle
	master.ReadTimeout = readTimeout
	instanceMaster := redisv9.NewClient(master)
	if err = instanceMaster.Ping(context.Background()).Err(); err != nil {
		panic("redis Unable to connect to RedisDefaultMaster," + err.Error())
	}

	if len(configRedisDefault.Slaves.Addr) == 0 {
		panic("redis RedisDefaultSlaves is empty")
	}

	instanceSlaves := make([]*redisv9.Client, 0)

	for _, addr := range configRedisDefault.Slaves.Addr {
		slave, err1 := redisv9.ParseURL(addr)
		if err1 != nil {
			panic("redis Unable to ParseURL to RedisDefaultSlaves," + err.Error())
		}
		slave.DialTimeout = dialTimeout
		slave.ConnMaxIdleTime = connMaxIdleTime
		slave.MaxActiveConns = maxActive
		slave.MaxIdleConns = maxIdle
		slave.ReadTimeout = readTimeout
		instanceSlave := redisv9.NewClient(slave)
		if err = instanceSlave.Ping(context.Background()).Err(); err != nil {
			panic("redis Unable to connect to RedisDefaultSlaves," + err.Error())
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
