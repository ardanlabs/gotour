//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This sample program shows you how to implement the sleeping barber
// problem.
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	const maxChairs = 3

	shop := OpenShop(maxChairs)
	defer shop.Close()

	// Close the shop in 50 milliseconds.
	t := time.NewTimer(50 * time.Millisecond)
	<-t.C
}

// =============================================================================

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
	chairs  chan customer  // The set of chairs customers wait in.
	wgClose sync.WaitGroup // Provides support for closing the shop.
}

// OpenShop creates a new shop for business and gets the barber working.
func OpenShop(maxChairs int) *Shop {
	fmt.Println("Opening the shop")

	s := Shop{
		chairs: make(chan customer, maxChairs),
	}
	atomic.StoreInt32(&s.open, 1)

	// This is the barber and they will service customers.

	s.wgClose.Add(1)
	go func() {
		defer s.wgClose.Done()

		fmt.Println("Barber ready to work")

		for cust := range s.chairs {
			s.serviceCustomer(cust)
		}
	}()

	// Start creating customers who enter the shop.

	go func() {
		var id int64

		for {
			// Wait some random time for the next customer to enter.
			time.Sleep(time.Duration(rand.Intn(75)) * time.Millisecond)

			name := fmt.Sprintf("cust-%d", atomic.AddInt64(&id, 1))
			if err := s.newCustomer(name); err != nil {
				if err == ErrShopClosed {
					break
				}
			}
		}
	}()

	return &s
}

// Close prevents any new customers from entering the shop and waits for
// the barber to finish all existing customers.
func (s *Shop) Close() {
	fmt.Println("Closing the shop")
	defer fmt.Println("Shop closed")

	// Mark the shop closed.
	atomic.StoreInt32(&s.open, 0)

	// Wait for the barber to finish with the existing customers.
	close(s.chairs)
	s.wgClose.Wait()
}

// =============================================================================

func (s *Shop) serviceCustomer(cust customer) {
	fmt.Printf("Barber servicing customer %q\n", cust.name)

	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)

	fmt.Printf("Barber finished customer %q\n", cust.name)

	if len(s.chairs) == 0 && atomic.LoadInt32(&s.open) == 1 {
		fmt.Println("Barber taking a nap")
	}
}

func (s *Shop) newCustomer(name string) error {
	if atomic.LoadInt32(&s.open) == 0 {
		fmt.Printf("Customer %q leaves, shop closed\n", name)
		return ErrShopClosed
	}

	fmt.Printf("Customer %q entered shop\n", name)

	select {
	case s.chairs <- customer{name: name}:
		fmt.Printf("Customer %q takes a seat and waits\n", name)

	default:
		fmt.Printf("Customer %q leaves, no seat\n", name)
	}

	return nil
}
