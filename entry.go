package main

import (
	"GoForIMServer/web"
	"flag"
)

var port string

func init() {
	flag.StringVar(&port, "port", "default", "8090")
}

func main() {
	flag.Parse()
	server := web.GetServerInstance(port)
	server.Run()
}
