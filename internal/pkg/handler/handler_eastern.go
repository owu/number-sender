package handler

import (
	"number-sender/internal/pkg/consts"
	"number-sender/internal/pkg/utils"
)

// EasternHandler , Taboo Numbers in Eastern Countries
type EasternHandler struct {
	BaseHandler
}

func (h *EasternHandler) Handle(number uint64) consts.Plans {
	if utils.Contains(number, 4) {
		return consts.Starter
	}
	return h.nextHandler(number)
}
