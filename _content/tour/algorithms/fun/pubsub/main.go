// This is a simple example put together to help a friend with the
// idea of not over-engineering a pubsub pattern. This is entirely
// production ready but more of a prototype and concept.
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	clients := NewClients()

	publisher := NewPublisher(clients)
	defer publisher.Shutdown()

	clients.Add("1")
	clients.Add("2")
	clients.Add("3")
	time.Sleep(time.Second)
	clients.Remove("2")
	time.Sleep(time.Second)
	clients.Remove("1")
	time.Sleep(time.Second)
	clients.Add("2")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	log.Println("shutting down")
}

// Publisher is consuming messages and publishing them.
type Publisher struct {
	clients *Clients
}

// NewPublisher connects to the publisher can recieves messages.
func NewPublisher(clients *Clients) *Publisher {
	pub := Publisher{
		clients: clients,
	}

	// This is not production ready since it can't be shutdown.
	// For now we just want a stream of messages.

	go func() {
		var counter int
		for {
			time.Sleep(100 * time.Millisecond)
			counter++
			log.Println("publisher: mesage received : sending to clients")
			clients.Send(fmt.Sprintf("message %d", counter))
		}
	}()

	return &pub
}

// Shutdown disconnects the publisher and stop messages.
func (p *Publisher) Shutdown() {

	// TO BE IMPLEMENTED
}

// Clients manage clients who are looking to receive messages.
type Clients struct {
	capacity int
	clients  map[string]chan string
	mu       sync.Mutex
}

// NewClients returns a clients management value.
func NewClients() *Clients {
	return &Clients{
		capacity: 1024,
		clients:  make(map[string]chan string),
	}
}

// Add places a client in the list for receiving messages.
func (c *Clients) Add(id string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.clients[id]; exists {
		return fmt.Errorf("client id already exists: %s", id)
	}

	c.clients[id] = make(chan string, c.capacity)

	return nil
}

// Remove takes a client out of the list.
func (c *Clients) Remove(id string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.clients[id]; !exists {
		return fmt.Errorf("client id doesn't exist: %s", id)
	}

	delete(c.clients, id)
	return nil
}

// Send will deliver the message to all existing clients.
func (c *Clients) Send(message string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// If the other side is not ready, they lose the message.
	// There is a capacity so messages would only be lost if
	// the client is not responding.

	for id, ch := range c.clients {
		select {
		case ch <- message:
			log.Println("client: sent: to: ", id)
		default:
			log.Println("client: timeout: to: ", id)
		}
	}
}
