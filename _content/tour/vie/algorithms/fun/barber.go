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
	fmt.Println("Mở cửa tiệm")

	s := Shop{
		chairs: make(chan customer, maxChairs),
	}
	atomic.StoreInt32(&s.open, 1)

	// This is the barber and they will service customers.

	s.wgClose.Add(1)
	go func() {
		defer s.wgClose.Done()

		fmt.Println("Thợ cắt tóc sẵn sàng làm việc")

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
	fmt.Println("Tiệm ngừng nhận khách")
	defer fmt.Println("Tiệm nghỉ")

	// Mark the shop closed.
	atomic.StoreInt32(&s.open, 0)

	// Wait for the barber to finish with the existing customers.
	close(s.chairs)
	s.wgClose.Wait()
}

// =============================================================================

func (s *Shop) serviceCustomer(cust customer) {
	fmt.Printf("Thợ cắt tóc phục vụ khách hàng %q\n", cust.name)

	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)

	fmt.Printf("Thợ cắt tóc cắt xong cho khách hàng %q\n", cust.name)

	if len(s.chairs) == 0 && atomic.LoadInt32(&s.open) == 1 {
		fmt.Println("Thợ cắt tóc ngủ trưa")
	}
}

func (s *Shop) newCustomer(name string) error {
	if atomic.LoadInt32(&s.open) == 0 {
		fmt.Printf("Khách hàng %q rời đi, tiệm nghỉ\n", name)
		return ErrShopClosed
	}

	fmt.Printf("Khách hàng %q bước vào tiệm\n", name)

	select {
	case s.chairs <- customer{name: name}:
		fmt.Printf("Khách hàng %q ngồi vào ghế chờ\n", name)

	default:
		fmt.Printf("Khách hàng %q rời đi, không có ghế chờ\n", name)
	}

	return nil
}
