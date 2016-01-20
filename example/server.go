package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/silalahi/go-sse"
)

var broker *sse.Broker

func main() {
	// SSE broker instance
	broker = sse.NewBroker()

	// Loop sending SSE message
	go func() {
		for i := 0; ; i++ {
			broker.Message <- fmt.Sprintf("%d - the time is %v", i, time.Now())
			log.Printf("Sent message %d ", i)
			time.Sleep(5 * 1e9)
		}
	}()

	// HTTP handler
	http.HandleFunc("/", home)
	http.HandleFunc("/event/", event)
	http.ListenAndServe(":8000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal("Unable to parse HTML template")
	}

	t.Execute(w, "Codigo")

	log.Println("Finished HTTP request at ", r.URL.Path)
}

func event(w http.ResponseWriter, r *http.Request) {
	f, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	// Create new channel of client for broker to send message
	messageChan := make(chan string)

	// Add channel to client map
	broker.NewClient <- messageChan

	// Listen to the closing of HTTP connection via CloseNotifier
	notify := w.(http.CloseNotifier).CloseNotify()
	go func() {
		<-notify
		// Removing closed client from client map
		broker.ClosedClient <- messageChan
		log.Println("HTTP connection just closed")
	}()

	// Set header to event streaming
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Listening
	for {
		// Read from message channel
		msg, open := <-messageChan

		if !open {
			// If connection closed, client has disconnected
			break
		}

		// Write to ResponseWriter
		fmt.Fprintf(w, "data: %s\n\n", msg)

		// Flushing
		f.Flush()
	}

	// Done
	log.Println("Finished HTTP request at ", r.URL.Path)

}
