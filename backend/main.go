package main

import (
	"To-Do-Go/backend/router"
	"net/http"
)

func main() {
	r := router.SetupRouter()

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
