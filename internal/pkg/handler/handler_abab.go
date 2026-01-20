package handler

import "number-sender/internal/pkg/consts"

// AbabHandler , the last 4 digits of abab are repeated
type AbabHandler struct {
	BaseHandler
}

func (h *AbabHandler) Handle(number uint64) consts.Plans {
	tail := number % 10000
	d1 := tail / 1000
	d2 := (tail / 100) % 10
	d3 := (tail / 10) % 10
	d4 := tail % 10

	if d1 == d3 && d2 == d4 {
		return consts.Premium
	}
	return h.nextHandler(number)
}
