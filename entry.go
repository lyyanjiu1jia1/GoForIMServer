package main

import (
	"GoForIMServer/web"
	"flag"
	"fmt"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8090", "the server port")
}

func main() {
	flag.Parse()
	fmt.Println("hello world")
	fmt.Println(port)
	server := web.GetServerInstance(port)
	server.Run()
}
