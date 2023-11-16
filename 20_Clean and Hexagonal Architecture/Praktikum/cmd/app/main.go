package main

import (
	r "capston-lms/internal/adapters/http"
)

func main() {
	e := r.InitRoutes()
	e.Debug = true
	e.Logger.Fatal(e.Start(":3000"))
}
