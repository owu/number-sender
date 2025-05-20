package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/owu/number-sender/internal/pkg/consts"
	"github.com/owu/number-sender/internal/pkg/logger"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
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
		panic("redis slave llen failed," + err.Error())
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
		if errors.Is(err, redis.Nil) {
			return 0
		}
		panic("redis master incrby failed," + err.Error())
	}
	return val
}

func (instance *DefaultRedis) LockObtain() bool {
	ret, err := instance.master().SetNX(context.Background(), instance.lockKey(), fmt.Sprint(time.Now().UnixMilli()), 60*time.Second).Result()
	if err != nil {
		return false
	}
	return ret
}

func (instance *DefaultRedis) LockRelease() {
	_, err := instance.master().Del(context.Background(), instance.lockKey()).Result()
	if err != nil {
		logger.Log.Error(err.Error())
	}
}

func (instance *DefaultRedis) push(plan consts.Plans, result []uint64) (int64, bool) {
	if len(result) == 0 {
		return 0, true
	}
	args := lo.Map(result, func(key uint64, _ int) string {
		return strconv.FormatUint(key, 10)
	})
	sort.Strings(args)
	ret, err := instance.master().RPush(context.Background(), instance.planKey(plan), args).Result()
	if err != nil {
		return 0, false
	}
	return ret, true
}

func (instance *DefaultRedis) PushMap(values map[consts.Plans][]uint64) {
	if len(values) == 0 {
		return
	}
	for plan, result := range values {
		if len(result) == 0 {
			continue
		}
		instance.push(plan, result)
	}
}
