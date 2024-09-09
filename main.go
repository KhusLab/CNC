package main

import (
	"github.com/wilsdane/C2Dev24/c2"
)

func main() {
	listener := c2.HttpListener{
		Ip:   "127.0.0.1",
		Port: "8080",
	}
	listener.Listen()
}
