//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος προκειμένου να παρουσιαστεί η σύνθεση struct.
package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

// Ο Data είναι η δομή των δεδομένων που θα αντιγράψουμε.
type Data struct {
	Line string
}

// =============================================================================

// Ο Xenia είναι ένα σύστημα από το οποίο θέλουμε να αντλήσουμε στοιχεία.
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

// Ο Pillar είναι ένα σύστημα που χρειαζόμαστε να αποθηκεύει δεδομένα.
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

// Ο System συνδυάζει τον Xenia και τον Pillar σε ένα σύστημα.
type System struct {
	Xenia
	Pillar
}

// =============================================================================

// Η pull γνωρίζει πως να αντλεί δεδομένα από τον Xenia.
func pull(x *Xenia, data []Data) (int, error) {
	for i := range data {
		if err := x.Pull(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

// Η store γνωρίζει πως να αποθηκεύει δεδομένα στον Pillar.
func store(p *Pillar, data []Data) (int, error) {
	for i := range data {
		if err := p.Store(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

// Η Copy γνωρίζει πως να αντλει και να αποθηκεύει γεγονότα στο σύστημα.
func Copy(sys *System, batch int) error {
	data := make([]Data, batch)

	for {
		i, err := pull(&sys.Xenia, data)
		if i > 0 {
			if _, err := store(&sys.Pillar, data[:i]); err != nil {
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
