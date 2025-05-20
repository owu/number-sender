package handler

import (
	"github.com/owu/number-sender/internal/pkg/consts"
	"strconv"
	"strings"
)

// RptAbHandler ,
// all the numbers are ab repeated
// all the numbers before are ab repeated and the last digit is a
type RptAbHandler struct {
	BaseHandler
}

func (h *RptAbHandler) Handle(number uint64) consts.Plans {
	str := strconv.FormatUint(number, 10)
	if len(str)%2 == 0 {
		//ab ab ab...
		if str == strings.Repeat(str[0:2], len(str)/2) {
			//fmt.Println(str)
			return consts.Ultimate
		}
	} else {
		//ab ab ab...a
		if str[0] == str[len(str)-1] && strings.HasPrefix(str, strings.Repeat(str[:2], (len(str)-1)/2)) {
			//fmt.Println(str)
			return consts.Ultimate
		}
	}
	return h.nextHandler(number)
}
