package handler

import (
	"github.com/owu/number-sender/internal/pkg/consts"
	"strconv"
)

// AaaaHandler , the last four digits are the same
type AaaaHandler struct {
	BaseHandler
}

func (h *AaaaHandler) Handle(number uint64) consts.Plans {
	if lastFourSame(number) {
		return consts.Ultimate
	}
	return h.nextHandler(number)
}

func lastFourSame(n uint64) bool {
	s := strconv.FormatUint(n, 10)
	if len(s) < 4 {
		return false
	}
	lastFour := s[len(s)-4:]

	for i := 1; i < 4; i++ {
		if lastFour[i] != lastFour[0] {
			return false
		}
	}
	return true
}
