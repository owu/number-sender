package handler

import "github.com/owu/number-sender/internal/pkg/consts"

type Handler interface {
	Handle(number uint64) consts.Plans
}

type BaseHandler struct {
	Next Handler
}

func (b *BaseHandler) nextHandler(number uint64) consts.Plans {
	if b.Next != nil {
		return b.Next.Handle(number)
	}
	return consts.Standard
}

// SetNext sets the next handler in the chain
func (b *BaseHandler) SetNext(next Handler) {
	b.Next = next
}
