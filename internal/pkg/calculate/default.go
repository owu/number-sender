package calculate

import (
	"github.com/owu/number-sender/internal/pkg/config"
	"github.com/owu/number-sender/internal/pkg/handler"
	"reflect"
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
		v := reflect.ValueOf(current).Elem()
		if base := v.FieldByName("BaseHandler"); base.IsValid() {
			if nextField := base.FieldByName("Next"); nextField.CanSet() {
				nextField.Set(reflect.ValueOf(next))
			}
		}
	}
	return handlers[0]
}
