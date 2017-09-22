package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/access"
	"github.com/go-ozzo/ozzo-routing/content"
	"github.com/go-ozzo/ozzo-routing/fault"
	"github.com/go-ozzo/ozzo-routing/slash"
)

func main() {
	router := routing.New()

	router.Use(
		// all these handlers are shared by every route
		access.Logger(log.Printf),
		slash.Remover(http.StatusMovedPermanently),
		fault.Recovery(log.Printf),
	)

	// serve RESTful APIs
	api := router.Group("/api")
	api.Use(
		// these handlers are shared by the routes in the api group only
		content.TypeNegotiator(content.JSON, content.XML),
	)
	api.Get("/users", func(c *routing.Context) error {
		return c.Write("user list")
	})
	api.Post("/users", func(c *routing.Context) error {
		return c.Write("create a new user")
	})
	api.Put(`/users/<id:\d+>`, func(c *routing.Context) error {
		return c.Write("update user " + c.Param("id"))
	})

	// Hello world.

	router.Get("/", func(c *routing.Context) error {
		return c.Write("Hello world")
	})

	// Streaming.

	router.Get("/streaming", func(c *routing.Context) error {
		// Observation: This buffers into large chunks before sending.

		// Need to cast http.ResponseWriter as a "flusher".

		w := c.Response
		flusher, ok := w.(http.Flusher)
		if !ok {
			panic("expected http.ResponseWriter to be an http.Flusher")
		}

		// Stream the response body.  Note: loop will not stop.

		ticker := time.NewTicker(time.Millisecond * 500)
		for aTime := range ticker.C {
			c.Write([]byte(aTime.String()))
			flusher.Flush()
		}

		return c.Write(">>> Shouldn't get here. <<<")
	})

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
