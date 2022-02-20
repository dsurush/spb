package main

import (
	"log"
	"spb/cmd/app"
	_ "spb/loginit"
)

func main() {
	log.Println("server init...")
	app.RouterInit()
}
