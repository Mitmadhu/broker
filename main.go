package main

import (
	"github.com/Mitmadhu/broker/server"
)

func main() {
	println("hello from broker")
	// err := model.User{}.Register("ayush","123", "ayush", "lasname", 18)
	// if err != nil {
	// 	println(err.Error())
	// }
	server.Routers()

}
