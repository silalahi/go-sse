package sse

import (
	"log"
)

// Broker responsible for keeping a list of clients (browsers) and broadcast event
type Broker struct {
	// Creating a map of client, keys are the channels which we can push message to client
	Clients map[chan string]bool
	// Channel of new client attached
	NewClient chan chan string
	// Channel of client that disconnected
	ClosedClient chan chan string
	// Channel of message to broadcast to clients
	Message chan string
}

// NewBroker return instance of broker
func NewBroker() (broker *Broker) {
	// Instance of broker
	broker = &Broker{
		make(map[chan string]bool),
		make(chan (chan string)),
		make(chan (chan string)),
		make(chan string),
	}

	// Run broker, listening and broadcasting events
	go broker.start()

	return broker
}

// Start a new goroutine that handle addition or removal of client,
// as well as broadcasting message to client that are currently attached
func (b *Broker) start() {
	go func() {
		// Loop endlessly
		for {
			// Block and switch type of event
			select {
			case s := <-b.NewClient:
				// New client attached
				b.Clients[s] = true
				log.Printf("New client added. %d registered clients", len(b.Clients))

			case s := <-b.ClosedClient:
				// Client disconnected
				delete(b.Clients, s)
				close(s)
				log.Printf("Removed client. %d registered clients", len(b.Clients))

			case msg := <-b.Message:
				// New message to send to each client
				for s := range b.Clients {
					s <- msg
				}
				log.Printf("Broadcase message to %d clients", len(b.Clients))
			}
		}
	}()
}
