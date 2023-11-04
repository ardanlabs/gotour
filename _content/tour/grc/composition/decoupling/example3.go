//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος, προκειμένου να παρουσιάσει την σύνθεση διεπαφών.
package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

// Ο Data είναι η δομή των δεομένων, που αντιγράφουμε.
type Data struct {
	Line string
}

// =============================================================================

// Ο Puller δηλώνει συμπεριφορά άντλησης δεδομένων.
type Puller interface {
	Pull(d *Data) error
}

// Ο Storer δηλώνει συμπεριφορά αποθήκευσης δεδομένων.
type Storer interface {
	Store(d *Data) error
}

// Ο PullStorer δηλώνει συμπεριφορά, τόσο για άντληση όσο και για αποθήκευση.
type PullStorer interface {
	Puller
	Storer
}

// =============================================================================

// Ο Xenia είναι ένα σύστημα, από το οποίο θέλουμε να αντλούμε δεδομένα.
type Xenia struct {
	Host    string
	Timeout time.Duration
}

// Η Pull γνωρίζει πως να αντλεί δεδομένα από τον Xenia.
func (*Xenia) Pull(d *Data) error {
	switch rand.Intn(10) {
	case 1, 9:
		return io.EOF

	case 5:
		return errors.New("error reading data from Xenia")

	default:
		d.Line = "Data"
		fmt.Println("In:", d.Line)
		return nil
	}
}

// Ο Pillar είναι ένα σύστημα, στο οποίο πρέπει να αποθηκεύουμε δεδομένα.
type Pillar struct {
	Host    string
	Timeout time.Duration
}

// Η Store γνωρίζει πως να αποθηκεύει δεδομένα στον Pillar.
func (*Pillar) Store(d *Data) error {
	fmt.Println("Out:", d.Line)
	return nil
}

// =============================================================================

// Ο System συνδυάζει έναν Xenia και έναν Pillar, σε ένα κοινό σύστημα.
type System struct {
	Xenia
	Pillar
}

// =============================================================================

// Η pull γνωρίζει πως να αντελι δεδομένα από κάθε Puller.
func pull(p Puller, data []Data) (int, error) {
	for i := range data {
		if err := p.Pull(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

// η store γνωρίζει πως να αποθηκεύει δεδομένα σε κάθε Storer.
func store(s Storer, data []Data) (int, error) {
	for i := range data {
		if err := s.Store(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

// Η Copy γνωρίζει πως να αντλεί και να αποθηκεύει δεδομένα, από κάθε σύστημα.
func Copy(ps PullStorer, batch int) error {
	data := make([]Data, batch)

	for {
		i, err := pull(ps, data)
		if i > 0 {
			if _, err := store(ps, data[:i]); err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}
	}
}

// =============================================================================

func main() {
	sys := System{
		Xenia: Xenia{
			Host:    "localhost:8000",
			Timeout: time.Second,
		},
		Pillar: Pillar{
			Host:    "localhost:9000",
			Timeout: time.Second,
		},
	}

	if err := Copy(&sys, 3); err != io.EOF {
		fmt.Println(err)
	}
}
