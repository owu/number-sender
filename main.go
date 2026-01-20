package main

import (
	"number-sender/internal/app/server"
)

func main() {
	if err := server.InitApp(); err != nil {
		panic("main server.InitApp failed," + err.Error())
	}
}
