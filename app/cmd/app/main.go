package main

import (
	"github.com/VusalShahbazov/chat-app/app/server"
	"log"
)

func main()  {
	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}
}


