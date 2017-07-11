package main

import (
	"net/http"

	"github.com/go-ozzo/ozzo-routing"
)

func main() {
	router := routing.New()

	router.Get("/", func(c *routing.Context) error {
		return c.Write("Hello world")
	})

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
