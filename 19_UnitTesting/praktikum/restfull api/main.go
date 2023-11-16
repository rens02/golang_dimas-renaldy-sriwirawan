package main

import (
	"code_structure/config"
	"code_structure/routes"
)



func main() {
	
	config.InitDB()
	e := routes.New()
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":3000"))
}