package main

import (
	"github.com/katsuomi/NotifyOfficeIlluminance/db"
	"github.com/katsuomi/NotifyOfficeIlluminance/server"
)

func main() {
	db.Init()
	server.Init()

	db.Close()
}
