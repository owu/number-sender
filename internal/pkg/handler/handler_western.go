package handler

import (
	"number-sender/internal/pkg/consts"
	"number-sender/internal/pkg/utils"
)

// WesternHandler , Taboo Numbers in Western Countries
type WesternHandler struct {
	BaseHandler
}

func (h *WesternHandler) Handle(number uint64) consts.Plans {
	if utils.Contains(number, 13) {
		return consts.Starter
	}
	return h.nextHandler(number)
}
