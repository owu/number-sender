package utils

var (
	Succeed          = Result{Msg: "success", Data: struct{}{}}
	Failed           = Result{Error: 10000, Msg: "Failed", Data: struct{}{}}
	MissingParams    = Result{Error: 10001, Msg: "Missing parameters", Data: struct{}{}}
	TimestampError   = Result{Error: 10002, Msg: "Timestamp error", Data: struct{}{}}
	TimestampExpired = Result{Error: 10003, Msg: "Timestamp expired", Data: struct{}{}}
	AuthFailed       = Result{Error: 10004, Msg: "Auth failed", Data: struct{}{}}
	LimitError       = Result{Error: 10005, Msg: "Out of rate limit", Data: struct{}{}}
)
