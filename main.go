package main

import (
	"TodoList/conf"
	"TodoList/routes"
)

func main() {
	conf.Init()
	routes.NewRouter().Run(conf.HttpPort)
}
