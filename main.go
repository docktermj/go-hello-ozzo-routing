package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-ozzo/ozzo-routing"
)

func main() {
	router := routing.New()

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
