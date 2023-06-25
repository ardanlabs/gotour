//go:build OMIT
// +build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Program creates customers for the simulation of the sleeping barber problem
// implemented in the shop package.
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"time"
)

// EnterCustomer is called to create a customer to be serviced. If
// the shop is closed the function returns an error. If the shop is open,
// a goroutine is created to handle the customers concurrently.
func (s *Shop) EnterCustomer(name string) error {
	if atomic.LoadInt32(&s.open) == 0 {
		return ErrShopClosed
	}

	s.wgEnter.Add(1)
	go func() {
		defer s.wgEnter.Done()
		select {
		case s.chairs <- customer{name: name}:
		default:
			fmt.Printf("No chair for customer %q\n", name)
		}
	}()

	return nil
}

func main() {
	const maxChairs = 10
	s := Open(maxChairs)

	// Create a goroutine that is constantly, but inconsistently, generating
	// customers who are entering the shop.
	go func() {
		var id int64
		for {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			name := fmt.Sprintf("cust-%d", atomic.AddInt64(&id, 1))
			if err := s.EnterCustomer(name); err != nil {
				fmt.Printf("Customer %q told %q\n", name, err)
				if err == ErrShopClosed {
					break
				}
			}
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	fmt.Println("Shutting down shop")
	s.Close()
}

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

var (
	// ErrShopClosed is returned when the shop is closed.
	ErrShopClosed = errors.New("shop closed")

	// ErrNoChair is returned when all the chairs are occupied.
	ErrNoChair = errors.New("no chair available")
)

// customer represents a customer to be serviced.
type customer struct {
	name string
}

// Shop represents the barber's shop which contains chairs for customers
// that customers can occupy and the barber can service. The shop can
// be closed for business.
type Shop struct {
	open    int32          // Determines if the shop is open for business.
	chairs  chan customer  // The set of chairs in the shop.
	wgClose sync.WaitGroup // Provides support for closing the shop.
	wgEnter sync.WaitGroup // Tracks customers entering the shop.
}

// Open creates a new shop for business and gets the barber working.
func Open(maxChairs int) *Shop {
	s := Shop{
		chairs: make(chan customer, maxChairs),
	}
	atomic.StoreInt32(&s.open, 1)

	// Get the barber working.
	s.wgClose.Add(1)
	go func() {
		defer s.wgClose.Done()
		for cust := range s.chairs {
			fmt.Printf("Barber servicing customer %q\n", cust.name)
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			fmt.Printf("Barber finished  customer %q\n", cust.name)
		}
	}()

	return &s
}

// Close prevents any new customers from entering the shop and waits for
// the barber to finish all existing customers.
func (s *Shop) Close() {

	// Mark the shop closed.
	atomic.StoreInt32(&s.open, 0)

	// Wait for an new customers just entering to be handled.
	s.wgEnter.Wait()

	// Wait for the barber to finish with the existing customers.
	close(s.chairs)
	s.wgClose.Wait()
}
