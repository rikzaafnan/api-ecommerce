package main

import (
	"api-ecommerce/app"
	"api-ecommerce/app/rest"
	"log"
)

func main() {
	log.Println("api-ecommerce")

	// initialize logging
	app.InitLogger()

	rest.StartApp()
}
