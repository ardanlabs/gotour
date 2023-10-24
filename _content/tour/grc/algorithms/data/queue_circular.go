//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Αυτό το δείγμα προγράμματος παρουσιάζει τον τρόπο συγγραφής μιας
// απλής κυκλικής ουράς.
package main

import (
	"errors"
	"fmt"
)

func main() {
	const items = 5

	q, err := New(items)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < items; i++ {
		name := fmt.Sprintf("Name%d", i)
		if err := q.Enqueue(Data{Name: name}); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("queue:", name)
	}

	fmt.Println("------------------")

	f := func(d Data) error {
		fmt.Println("enqueue:", d.Name)
		return nil
	}

	q.Operate(f)
}

// Ο Data αναπαριστά αυτό που θα αποθηκευτεί στην ουρά.
type Data struct {
	Name string
}

// Ο Queue αναπαριστά μια λίστα δεδομένων.
type Queue struct {
	Count int
	data  []Data
	front int
	end   int
}

// Η New επιστρέφει μια ουρά, με δεδομένη χωρητικότητα.
func New(cap int) (*Queue, error) {
	if cap <= 0 {
		return nil, errors.New("invalid capacity")
	}

	q := Queue{
		front: 0,
		end:   0,
		data:  make([]Data, cap),
	}
	return &q, nil
}

// Η Enqueue εισάγει δεδομένα στην ουρά, εφόσον
// υπάρχει διαθέσιμη χωρητικότητα.
func (q *Queue) Enqueue(data Data) error {

	// Αν το μπροστινό μέρος της ουράς είναι ακριβώς από
	// πίσω, από το τελευταίο μέρος της ή αν το μπροστινό μέρος
	// βρίσκεται στο τέλος της χωρητικότητας και το τελευταίο μέρος
	// είναι στην αρχή της χωρητικότητας, τότε η ουρά είναι
	// γεμάτη.
	//
	//  F  E  - Enqueue (Full) |  E        F - Enqueue (Full)
	// [A][B][C]               | [A][B][C]
	if q.front+1 == q.end ||
		q.front == len(q.data) && q.end == 0 {
		return errors.New("queue at capacity")
	}

	switch {
	case q.front == len(q.data):

		// Αν βρισκόμαστε στο τέλος της χωρητικότητας, τότε
		// μετακινούμαστε στην αρχή της χωρητικότητας,
		// μετακινώντας τον μπροστινό δείκτη διεύθυνσης,
		// στην αρχή.
		q.front = 0
		q.data[q.front] = data
	default:

		// Προσθέστε τα δεδομένα στην τρέχουσα μπροστινή
		// θέση και στην συνέχεια μετακινείστε τον μπροστινό
		// δείκτη διεύθυνσης.
		q.data[q.front] = data
		q.front++
	}

	q.Count++

	return nil
}

// Η Dequeue απομακρύνει δεδομένα από την ουρά, αν αυτά
// υπάρχουν ήδη, μέσα.
func (q *Queue) Dequeue() (Data, error) {

	// Αν το μπροστινό και το τελευταίο μέρος είναι τα ίδια,
	// τότε η ουρά είναι άδεια.
	//
	//  EF - (Empty)
	// [  ][ ][ ]
	if q.front == q.end {
		return Data{}, errors.New("queue is empty")
	}

	var data Data
	switch {
	case q.end == len(q.data):

		// Αν βρισκόμαστε στο τέλος της χωρητικότητας,
		// τότε μετακινούμαστε στην αρχή της χωρητικότητας,
		// μετακινώντας τον δείκτη διεύθυνσης τέλους,
		// στην αρχή.
		q.end = 0
		data = q.data[q.end]
	default:

		// Απομακρύνετε τα δεδομένα από την τρέχουσα θέση τέλους
		// και μετά μετακινείστε τον δείκτη διεύθυνσης τέλους.
		data = q.data[q.end]
		q.end++
	}

	q.Count--

	return data, nil
}

// Η Operate δέχεται μια συνάρτηση, που παίρνει δεδομένα
// και καλεί την συγκεκριμένη συνάρτηση για κάθε
// κομμάτι δεδομένων, που μπορεί να βρει.
func (q *Queue) Operate(f func(d Data) error) error {
	end := q.end
	for {
		if end == q.front {
			break
		}

		if end == len(q.data) {
			end = 0
		}

		if err := f(q.data[end]); err != nil {
			return err
		}

		end++
	}
	return nil
}
