package handler

import (
	"github.com/owu/number-sender/internal/pkg/consts"
	"strconv"
)

// DcbaHandler , The last four digits decrease
type DcbaHandler struct {
	BaseHandler
}

func (h *DcbaHandler) Handle(number uint64) consts.Plans {
	if decreasing(number) {
		return consts.Ultimate
	}
	return h.nextHandler(number)
}

func decreasing(num uint64) bool {
	s := strconv.Itoa(int(num))
	if len(s) < 4 {
		return false
	}

	lastFour := s[len(s)-4:]
	digits := make([]int, 4)
	for i := 0; i < 4; i++ {
		digits[i] = int(lastFour[i] - '0')
	}

	isConsecutiveDecreasing := true
	for i := 0; i < 3; i++ {
		if digits[i]-digits[i+1] != 1 {
			isConsecutiveDecreasing = false
			break
		}
	}

	return isConsecutiveDecreasing
}
