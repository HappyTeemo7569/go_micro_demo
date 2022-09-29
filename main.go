package main

import "go_micro_demo/server"

func main() {
	server.RunHttpByGin()
	server.RunHttpByMicro()
}
