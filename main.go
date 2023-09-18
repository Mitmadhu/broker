package main

import (
	"github.com/Mitmadhu/broker/server"
	"github.com/Mitmadhu/mysqlDB/config"
)



func main(){
	println(config.Cnf.DB, config.Cnf.DB == nil)
	println("hello from broker")
	server.Routers()
}