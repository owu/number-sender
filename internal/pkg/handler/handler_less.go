package handler

import "github.com/owu/number-sender/internal/pkg/consts"

// LessHandler , filter out numbers less than 10000
type LessHandler struct {
	BaseHandler
}

func (h *LessHandler) Handle(number uint64) consts.Plans {
	if number <= consts.Less {
		return consts.Reserved
	}
	return h.nextHandler(number)
}
