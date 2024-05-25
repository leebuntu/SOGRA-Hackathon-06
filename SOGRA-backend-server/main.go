package main

import (
	"SOGRA/db_controller"
	"SOGRA/router"
)

func main() {
	db_controller.InitDB() // Init database
	router.InitServer()    // Init gin router
}
