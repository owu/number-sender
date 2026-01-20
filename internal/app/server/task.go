package server

import (
	"number-sender/internal/pkg/calculate"
	"number-sender/internal/pkg/consts"
	"number-sender/internal/pkg/redis"
)

func Task(defaultRedis *redis.DefaultRedis, chains *calculate.Chains) {

	if !defaultRedis.LockObtain() {
		return
	}
	defer defaultRedis.LockRelease()

	length := defaultRedis.Len(consts.Standard)
	if length >= consts.StandardMax {
		return
	}

	radix := defaultRedis.IncrRadix(consts.Batch)

	result := make(map[consts.Plans][]uint64)
	result[consts.Starter], result[consts.Standard], result[consts.Premium], result[consts.Ultimate] = chains.Decide(radix-consts.Batch+1, radix)

	defaultRedis.PushMap(result)
}
