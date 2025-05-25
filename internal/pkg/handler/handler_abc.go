package handler

import "github.com/owu/number-sender/internal/pkg/consts"

// AbcHandler , The last three digits increase
type AbcHandler struct {
	BaseHandler
}

func (h *AbcHandler) Handle(number uint64) consts.Plans {
	tail := number % 1000
	a := tail / 100
	b := (tail / 10) % 10
	c := tail % 10
	if b == a+1 && c == b+1 {
		return consts.Premium
	}
	return h.nextHandler(number)
}
