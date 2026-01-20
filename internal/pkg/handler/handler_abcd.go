package handler

import (
	"number-sender/internal/pkg/consts"
	"strconv"
)

// AbcdHandler , The last four digits increase
type AbcdHandler struct {
	BaseHandler
}

func (h *AbcdHandler) Handle(number uint64) consts.Plans {
	if increasing(number) {
		return consts.Ultimate
	}
	return h.nextHandler(number)
}

func increasing(num uint64) bool {
	s := strconv.Itoa(int(num))
	if len(s) < 4 {
		return false
	}
	lastFour := s[len(s)-4:]
	digits := make([]int, 4)
	for i := 0; i < 4; i++ {
		digits[i] = int(lastFour[i] - '0')
	}
	isConsecutiveIncreasing := true
	for i := 0; i < 3; i++ {
		if digits[i+1]-digits[i] != 1 {
			isConsecutiveIncreasing = false
			break
		}
	}
	return isConsecutiveIncreasing
}
