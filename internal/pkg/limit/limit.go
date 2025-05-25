package limit

import (
	"github.com/juju/ratelimit"
	"github.com/owu/number-sender/internal/pkg/consts"
)

func (instance *Limiter) TakeAvailable(fullPath string) bool {
	bucket, status := instance.BucketMap.Load(fullPath)
	if status && bucket != nil {
		return bucket.(*ratelimit.Bucket).TakeAvailable(1) == 1
	}
	newBucket := ratelimit.NewBucketWithRate(float64(consts.RateLimit), consts.RateLimit)

	instance.BucketMap.Store(fullPath, newBucket)
	return newBucket.TakeAvailable(1) == 1
}
