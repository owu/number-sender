package handler

import (
	"number-sender/internal/pkg/consts"
	"strconv"
)

// TwoHandler , it consists of only two kinds of numbers
type TwoHandler struct {
	BaseHandler
}

func (h *TwoHandler) Handle(number uint64) consts.Plans {
	if hasOnlyTwoDigits(number) {
		return consts.Ultimate
	}
	return h.nextHandler(number)
}

func hasOnlyTwoDigits(n uint64) bool {
	if n <= 100 {
		return false
	}
	s := strconv.FormatUint(n, 10)
	var digits [10]bool
	for _, c := range s {
		digits[c-'0'] = true
	}
	count := 0
	for _, exists := range digits {
		if exists {
			count++
		}
	}
	return count == 2
}
