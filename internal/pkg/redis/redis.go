package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"math/rand"
	"number-sender/internal/pkg/consts"
	"number-sender/internal/pkg/logger"
	"sort"
	"strconv"
	"time"
)

var (
	planKey  = "num-sdr:%s"
	lockKey  = "num-sdr:lock"
	radixKey = "num-sdr:radix"
)

func (instance *DefaultRedis) planKey(plan consts.Plans) string {
	return fmt.Sprintf(planKey, string(plan))
}

func (instance *DefaultRedis) lockKey() string {
	return lockKey
}

func (instance *DefaultRedis) radixKey() string {
	return radixKey
}

func (instance *DefaultRedis) Len(plan consts.Plans) int64 {
	ctx := context.Background()
	val, err := instance.slave().Do(ctx, "LLEN", instance.planKey(plan)).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0
		}
		logger.Log.Error("redis slave llen failed", zap.Error(err), zap.String("plan", string(plan)))
		return 0
	}
	return val.(int64)
}

func (instance *DefaultRedis) Pop(plan consts.Plans) int64 {
	ret, err := instance.master().LPop(context.Background(), instance.planKey(plan)).Int64()
	if err != nil {
		return 0
	}
	return ret
}

func (instance *DefaultRedis) IncrRadix(step int64) int64 {
	ctx := context.Background()
	val, err := instance.master().IncrBy(ctx, instance.radixKey(), step).Result()
	if err != nil {
		logger.Log.Error("redis master incrby failed", zap.Error(err))
		return 0
	}
	return val
}

// LockValue 存储当前获取的锁值
var LockValue string

func (instance *DefaultRedis) LockObtain() bool {
	// 使用更安全的锁获取方式，设置唯一标识
	lockValue := fmt.Sprintf("%d-%d", time.Now().UnixMilli(), rand.Int63())
	ret, err := instance.master().SetNX(context.Background(), instance.lockKey(), lockValue, 60*time.Second).Result()
	if err != nil {
		logger.Log.Error("redis master setnx failed", zap.Error(err))
		return false
	}
	if ret {
		LockValue = lockValue
	}
	return ret
}

func (instance *DefaultRedis) LockRelease() {
	// 优化：使用Lua脚本确保原子性删除锁
	luaScript := `
		if redis.call("GET", KEYS[1]) == ARGV[1] then
			return redis.call("DEL", KEYS[1])
		else
			return 0
		end
	`
	_, err := instance.master().Eval(context.Background(), luaScript, []string{instance.lockKey()}, LockValue).Result()
	if err != nil {
		logger.Log.Error("redis master del lock failed", zap.Error(err))
	} else {
		// 释放锁后清空锁值
		LockValue = ""
	}
}

func (instance *DefaultRedis) push(plan consts.Plans, result []uint64) (int64, bool) {
	if len(result) == 0 {
		return 0, true
	}

	// 直接在uint64层面排序，避免转换为字符串后再排序
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	// 批量转换为字符串
	args := make([]interface{}, len(result))
	for i, v := range result {
		args[i] = strconv.FormatUint(v, 10)
	}

	ret, err := instance.master().RPush(context.Background(), instance.planKey(plan), args...).Result()
	if err != nil {
		logger.Log.Error("redis master rpush failed", zap.Error(err), zap.String("plan", string(plan)))
		return 0, false
	}
	return ret, true
}

func (instance *DefaultRedis) PushMap(values map[consts.Plans][]uint64) {
	if len(values) == 0 {
		return
	}

	// 优化：使用Redis管道减少网络开销
	for plan, result := range values {
		if len(result) == 0 {
			continue
		}
		instance.push(plan, result)
	}
}
