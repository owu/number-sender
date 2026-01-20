package calculate

import (
	"number-sender/internal/pkg/config"
	"number-sender/internal/pkg/handler"
)

type Chains struct {
	chains handler.Handler
}

func NewChains(configs *config.LoadConfigs) *Chains {
	var handlers []handler.Handler
	rules := configs.ApiRules()

	if rules.Less {
		handlers = append(handlers, &handler.LessHandler{})
	}

	if rules.Eastern {
		handlers = append(handlers, &handler.EasternHandler{})
	}

	if rules.Western {
		handlers = append(handlers, &handler.WesternHandler{})
	}

	if rules.Abcd {
		handlers = append(handlers, &handler.AbcdHandler{})
	}

	if rules.Abc {
		handlers = append(handlers, &handler.AbcHandler{})
	}

	if rules.Dcba {
		handlers = append(handlers, &handler.DcbaHandler{})
	}

	if rules.Cba {
		handlers = append(handlers, &handler.CbaHandler{})
	}

	if rules.Two {
		handlers = append(handlers, &handler.TwoHandler{})
	}

	if rules.Aaaa {
		handlers = append(handlers, &handler.AaaaHandler{})
	}

	if rules.Aaa {
		handlers = append(handlers, &handler.AaaHandler{})
	}

	if rules.RptAb {
		handlers = append(handlers, &handler.RptAbHandler{})
	}

	if rules.Abab {
		handlers = append(handlers, &handler.AbabHandler{})
	}

	handler := chains(handlers...)

	return &Chains{
		chains: handler,
	}
}

func chains(handlers ...handler.Handler) handler.Handler {
	for i := 0; i < len(handlers)-1; i++ {
		current := handlers[i]
		next := handlers[i+1]
		// 使用类型断言获取BaseHandler并设置Next，避免反射
		if baseHandler, ok := current.(interface{ SetNext(handler.Handler) }); ok {
			baseHandler.SetNext(next)
		} else {
			// 如果处理器没有SetNext方法，尝试直接设置Next字段
			if withBase, ok := current.(interface{ GetBase() *handler.BaseHandler }); ok {
				withBase.GetBase().Next = next
			}
		}
	}
	return handlers[0]
}
