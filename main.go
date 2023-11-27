package main

import (
	"github.com/Mitmadhu/broker/config"
	"github.com/Mitmadhu/broker/server"
)

func main() {
	println("hello from broker")
	config.Init()
	server.InitRouter()

}
