package model

type Fetch struct {
	Starter  int64 `json:"starter"`
	Standard int64 `json:"standard"`
	Premium  int64 `json:"premium"`
	Ultimate int64 `json:"ultimate"`
}
