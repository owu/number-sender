package calculate

import (
	"fmt"
	"github.com/owu/number-sender/internal/pkg/consts"
	"github.com/owu/number-sender/internal/pkg/logger"
	"go.uber.org/zap"
	"time"
)

func (instance *Chains) Decide(min, max int64) (starter, standard, premium, ultimate []uint64) {
	reserved := make([]uint64, 0)

	startT := time.Now()

	for i := uint64(min); i <= uint64(max); i++ {
		result := instance.chains.Handle(i)
		switch result {
		case consts.Reserved:
			reserved = append(reserved, i)
			continue
		case consts.Starter:
			starter = append(starter, i)
			continue
		case consts.Standard:
			standard = append(standard, i)
			continue
		case consts.Premium:
			premium = append(premium, i)
			continue
		case consts.Ultimate:
			ultimate = append(ultimate, i)
			continue
		default:
		}
	}
	logger.Log.Info("calculate.decide",
		zap.Int64("min", min),
		zap.Int64("max", max),
		zap.String("spent", fmt.Sprintf(":%fs, ", time.Since(startT).Seconds())),
		zap.String("len", fmt.Sprintf("%d,%d,%d,%d,%d", len(reserved), len(starter), len(standard), len(premium), len(ultimate))),
	)
	return
}
