package handler

import "github.com/owu/number-sender/internal/pkg/consts"

// AaaHandler , the last three digits are the same
type AaaHandler struct {
	BaseHandler
}

func (h *AaaHandler) Handle(number uint64) consts.Plans {
	if lastThreeSame(number) {
		return consts.Premium
	}
	return h.nextHandler(number)
}

func lastThreeSame(num uint64) bool {
	n := num
	if n < 100 {
		return false
	}
	lastThree := n % 1000
	d1 := lastThree / 100
	d2 := (lastThree / 10) % 10
	d3 := lastThree % 10
	return d1 == d2 && d2 == d3
}
