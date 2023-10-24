//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος, που παρουσιάζει πως επιδεικνύεται περισσότερη ακρίβεια
// στον σχεδιασμό API.
package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

// Ο Data είναι η δομή των δεδομένων, που αντιγράφουμε.
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

// =============================================================================

// Ο Xenia είναι ένα σύστημα, από το οποίο χρειάζεται να αντλήσουμε δεδομένα.
type Xenia struct {
	Host    string
	Timeout time.Duration
}

// Η Pull γνωρίζει, πως να αντλεί δεδομένα από τον Xenia.
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

// Ο Pillar είναι ένα σύστημα, στο οποίο χρειάζεται να αποθηκεύσουμε δεδομένα.
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

// Η pull γνωρίζει πως να αντλεί δεδομένα, από κάθε Puller.
func pull(p Puller, data []Data) (int, error) {
	for i := range data {
		if err := p.Pull(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

// Η store γνωρίζει πως να αποθηκεύει δεδομένα, σε κάθε Storer.
func store(s Storer, data []Data) (int, error) {
	for i := range data {
		if err := s.Store(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

// Η Copy γνωρίζει πως να αντλεί και να αποθηκεύει δεδομένα, από κάθε σύστημα.
func Copy(p Puller, s Storer, batch int) error {
	data := make([]Data, batch)

	for {
		i, err := pull(p, data)
		if i > 0 {
			if _, err := store(s, data[:i]); err != nil {
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
	x := Xenia{
		Host:    "localhost:8000",
		Timeout: time.Second,
	}

	p := Pillar{
		Host:    "localhost:9000",
		Timeout: time.Second,
	}

	if err := Copy(&x, &p, 3); err != io.EOF {
		fmt.Println(err)
	}
}
