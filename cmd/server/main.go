package main

import (

	"github.com/chetan177/oneclick/rest"
)

func main() {
	server := rest.NewServer()
	server.Start()
}
