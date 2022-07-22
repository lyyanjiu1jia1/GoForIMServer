package main

import "GoForIMServer/web"

func main() {
	server := web.GetServerInstance()
	server.Run()
}
