//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Αυτό το δείγμα προγράμματος παρουσιάζει πως μπορεί κάποιος να γράψει μια
// απλή διπλά συνδεδεμένη λίστα.
package main

import (
	"fmt"
	"strings"
)

func main() {
	var l List

	const nodes = 5
	for i := 0; i < nodes; i++ {
		data := fmt.Sprintf("Node%d", i)
		l.Add(data)
	}

	f := func(n *Node) error {
		fmt.Println("Data:", n.Data)
		return nil
	}

	l.Operate(f)

	fmt.Println("------------------")

	l.OperateReverse(f)
}

// =============================================================================

// Ο Node αναπαριστά τα δεδομένα προς αποθήκευση.
type Node struct {
	Data string
	next *Node
	prev *Node
}

// Η List αναπαριστά μια λίστα από στοιχεία (node).
type List struct {
	Count int
	first *Node
	last  *Node
}

// Η Add τοποθετεί ένα νέο στοιχείο στο τέλος της λίστας.
func (l *List) Add(data string) *Node {

	// Όταν κατασκευάζετε το νέο στοιχείο, βάλτε το νέο στοιχείο
	// να δείχνει προς το τελευταίο στοιχείο της λίστας.
	n := Node{
		Data: data,
		prev: l.last,
	}

	// Αυξήστε την μέτρηση για το νέο στοιχείο.
	l.Count++

	// Αν πρόκειται για το πρώτο στοιχείο,
	// τότε συνδέστε το.
	if l.first == nil && l.last == nil {
		l.first = &n
		l.last = &n
		return &n
	}

	// Διορθώστε το γεγονός ότι το τελευταίο στοιχείο δεν δείχνει
	// στο ΝΕΟ στοιχείο.
	l.last.next = &n

	// Διορθώστε το γεγονός ότι ο τελευταίος δείκτης διεύθυνσης δεν
	// δείχνει προς το πραγματικό τέλος της λίστας.
	l.last = &n

	return &n
}

// Η AddFront τοποθετεί ένα νέο στοιχείο στην αρχή της λίστας.
func (l *List) AddFront(data string) *Node {

	// Όταν δημιουργείτε ένα νέο στοιχείο, βάλτε το νέο στοιχείο
	// να δείχνει στο πρώτο στοιχείο της λίστας.
	n := Node{
		Data: data,
		next: l.first,
	}

	// Αυξήστε την μέτρηση για το νέο στοιχείο.
	l.Count++

	// Αν πρόκειται για το πρώτο στοιχείο,
	// τότε συνδέστε το.
	if l.first == nil && l.last == nil {
		l.first = &n
		l.last = &n
		return &n
	}

	// Διορθώστε το γεγονός ότι το πρώτο στοιχείο δεν δείχνει
	// στο ΝΕΟ στοιχείο.
	l.first.prev = &n

	// Διορθώστε το γεγονός ότι ο πρώτος δείκτης διεύθυνσης δεν
	// δείχνει προς την πραγματική αρχή της λίστας.
	l.first = &n

	return &n
}

// Η Find διατρέχει την λίστα αναζητώντας τα συγκεκριμένα δεδομένα.
func (l *List) Find(data string) (*Node, error) {
	n := l.first
	for n != nil {
		if n.Data == data {
			return n, nil
		}
		n = n.next
	}
	return nil, fmt.Errorf("unable to locate %q in list", data)
}

// Η FindReverse διατρέχει την λίστα στην αντίθετη κατεύθυνση,
// αναζητώντας τα συγκεκριμένα δεδομένα.
func (l *List) FindReverse(data string) (*Node, error) {
	n := l.last
	for n != nil {
		if n.Data == data {
			return n, nil
		}
		n = n.prev
	}
	return nil, fmt.Errorf("unable to locate %q in list", data)
}

// Η Remove διατρέχει την λίστα αναζητώντας για τα συγκεκριμένα
// δεδομένα και αν τα βρει, τα απομακρύνει από την λίστα.
func (l *List) Remove(data string) (*Node, error) {
	n, err := l.Find(data)
	if err != nil {
		return nil, err
	}

	// Αποκολλήστε το στοιχείο συνδέοντας τον δείκτη διεύθυνσης
	// next του προηγούμενου στοιχείου με το στοιχείο εμπρός από
	// εκείνο που απομακρύνετε.
	n.prev.next = n.next
	n.next.prev = n.prev
	l.Count--

	return n, nil
}

// Η Operate δέχεται μια συνάρτηση, που παίρνει ένα στοιχείο
// και καλεί την συγκεκριμένη συνάρτηση για κάθε στοιχείο.
func (l *List) Operate(f func(n *Node) error) error {
	n := l.first
	for n != nil {
		if err := f(n); err != nil {
			return err
		}
		n = n.next
	}
	return nil
}

// Η OperateReverse δέχεται μια συνάρτηση, που παίρνει ένα στοιχείο
// και καλεί την συγκεκριμένη συνάρτηση για κάθε στοιχείο.
func (l *List) OperateReverse(f func(n *Node) error) error {
	n := l.last
	for n != nil {
		if err := f(n); err != nil {
			return err
		}
		n = n.prev
	}
	return nil
}

// Η AddSort προσθέτει ένα στοιχείο με βάση μια λεξικογραφική
// ταξινόμηση.
func (l *List) AddSort(data string) *Node {

	// Αν η λίστα ήταν άδεια, προσθέστε τα δεδομένα
	// ως το πρώτο στοιχείο.
	if l.first == nil {
		return l.Add(data)
	}

	// Διατρέξτε την λίστα, αναζητώντας σημείο τοποθέτησης.
	n := l.first
	for n != nil {

		// Αν αυτά τα δεδομένα είναι μεγαλύτερα από το
		// τρέχον στοιχείο, συνεχίστε την προσπέλαση μέχρι
		// να είναι μικρότερα ή ίσα.
		if strings.Compare(data, n.Data) > 0 {
			n = n.next
			continue
		}

		// Δημιουργείστε ένα νέο στοιχείο και τοποθετήστε το
		// πριν από το τρέχον στοιχείο.
		new := Node{
			Data: data,
			next: n,
			prev: n.prev,
		}

		l.Count++

		// Αν αυτό το στοιχείο πρόκειται να είναι
		// τώρα το πρώτο, διορθώστε τον πρώτο δείκτη
		// διεύθυνσης.
		if l.first == n {
			l.first = &new
		}

		// Αν το τρέχον στοιχείο δείχνει σε προηγούμενο
		// στοιχείο, τότε το πεδίο next αυτού του προηγούμενου
		// στοιχείου, πρέπει να δείχνει στο νέο στοιχείο.
		if n.prev != nil {
			n.prev.next = &new
		}

		// Ο τρέχων δείκτης διεύθυνσης prev πρέπει να δείχνει
		// σε αυτό το νέο στοιχείο.
		n.prev = &new

		return n
	}

	// Αυτή πρέπει αν είναι η μεγαλύτερη συμβολοσειρά, επομένως
	// προσθέστε την στο τέλος.
	return l.Add(data)
}
