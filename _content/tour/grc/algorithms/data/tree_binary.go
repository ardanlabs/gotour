//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Αυτό το δείγμα προγράμματος παρουσιάζει τον τρόπο συγγραφής ενός
// βασικού δυαδικού δέντρου.
package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	values := []Data{
		{Key: 65, Name: "Bill"},
		{Key: 45, Name: "Ale"},
		{Key: 35, Name: "Joan"},
		{Key: 75, Name: "Hanna"},
		{Key: 85, Name: "John"},
		{Key: 78, Name: "Steph"},
		{Key: 95, Name: "Sally"},
	}

	var tree Tree
	for _, value := range values {
		tree.Insert(value)
	}

	PrettyPrint(tree)
	pre := tree.PreOrder()
	fmt.Println("Pre-order :", pre)
	in := tree.InOrder()
	fmt.Println("In-order  :", in)
	post := tree.PostOrder()
	fmt.Println("Post-order:", post)

	fmt.Print("\n")
	d35, err := tree.Find(35)
	if err != nil {
		fmt.Println("ERROR: Unable to find 35")
		return
	}
	fmt.Println("found:", d35)

	d78, err := tree.Find(78)
	if err != nil {
		fmt.Println("ERROR: Unable to find 78")
		return
	}
	fmt.Println("found:", d78)

	d3, err := tree.Find(3)
	if err == nil {
		fmt.Println("ERROR: found 3", d3)
		return
	}
	fmt.Println("not-found: 3")

	fmt.Print("\n")
	tree.Delete(75)
	PrettyPrint(tree)

	tree.Delete(85)
	PrettyPrint(tree)
}

// =============================================================================

// Ο Data αναπαριστά την πληροφορία προς αποθήκευση.
type Data struct {
	Key  int
	Name string
}

// Ο Tree αναπαριστά όλες τις τιμές στο δέντρο.
type Tree struct {
	root *node
}

// Η Insert προσθέτει μια τιμή στο δέντρο και κρατάει το δέντρο
// ισορροπημένο.
func (t *Tree) Insert(data Data) {
	t.root = t.root.insert(t, data)

	if t.root.balRatio() < -1 || t.root.balRatio() > 1 {
		t.root = t.root.rebalance()
	}
}

// Η Find διατρέχει το δέντρο, αναζητώντας για το συγκεκριμένο δέντρο.
func (t *Tree) Find(key int) (Data, error) {
	if t.root == nil {
		return Data{}, errors.New("cannot find from an empty tree")
	}

	return t.root.find(key)
}

// Η Delete απομακρύνει το κλειδί από το δέντρο και το
// κρατάει ισορροπημένο.
func (t *Tree) Delete(key int) error {
	if t.root == nil {
		return errors.New("cannot delete from an empty tree")
	}

	fakeParent := &node{right: t.root}
	if err := t.root.delete(key, fakeParent); err != nil {
		return err
	}

	if fakeParent.right == nil {
		t.root = nil
	}
	return nil
}

// Η διέλευση PreOrder παίρνει το στοιχείο της ρίζας και στην συνέχεια
// διασχίζει τα παιδιά της με αναδρομή.
// Περιπτώσεις χρήσης: αντιγραφή δέντρων, απεικόνιση συμβολισμού
// προθέματος (στμ. prefix notation).
//
//	      #1
//	   /      \
//	  #2      #5
//	 /  \    /  \
//	#3  #4  #6  #7
func (t *Tree) PreOrder() []Data {
	order := []Data{}
	f := func(n *node) {
		order = append(order, n.data)
	}
	t.root.preOrder(f)
	return order
}

// Η διέλευση InOrder επισκέπτεται από το πιο αριστερό στοιχείο
// στο πιο δεξιό στοιχείο, ανεξάρτητα από το βάθος.
// Η διέλευση In-order δίνει τις τιμές των στοιχείων,
// σε αύξουσα κατάταξη.
//
//	      #4
//	   /      \
//	  #2      #6
//	 /  \    /  \
//	#1  #3  #5  #7
func (t *Tree) InOrder() []Data {
	order := []Data{}
	f := func(n *node) {
		order = append(order, n.data)
	}
	t.root.inOrder(f)
	return order
}

// Η διέλευση PostOrder παίρνει το πιο αριστερό στοιχείο, στην συνέχεια
// επισκέπτεται το αδελφό στοιχείο και στην συνέχεια επισκέπτεται το
// πατρικό στοιχείο, αναδρομικά.
// Περιπτώσεις χρήσης: διαγραφή δέντρου, απεικόνιση συμβολισμού
// επιθέματος (στμ. postfix notation).
//
//	      #7
//	   /      \
//	  #3      #6
//	 /  \    /  \
//	#1  #2  #4  #5
func (t *Tree) PostOrder() []Data {
	order := []Data{}
	f := func(n *node) {
		order = append(order, n.data)
	}
	t.root.postOrder(f)
	return order
}

// =============================================================================

// Ο node αναπαριστά τα δεδομένα, που είναι αποθηκευμένα στο δέντρο.
type node struct {
	data  Data
	level int
	tree  *Tree
	left  *node
	right *node
}

// Η height επιστρέφει το επίπεδο του δέντρου στο οποίο υπάρχει
// το στοιχείο. Το επίπεδο 1 βρίσκεται στο τελευταίο επίπεδο του
// δέντρου.
//
//	      #7          -- height = 3
//	   /      \
//	  #3      #6      -- height = 2
//	 /  \    /  \
//	#1  #2  #4  #5    -- height = 1
func (n *node) height() int {
	if n == nil {
		return 0
	}
	return n.level
}

// Η insert προσθέτει το στοιχείο στο δέντρο και
// επιβεβαιώνει ότι το δέντρο παραμένει ισορροπημένο.
func (n *node) insert(t *Tree, data Data) *node {
	if n == nil {
		return &node{data: data, level: 1, tree: t}
	}

	switch {
	case data.Key < n.data.Key:
		n.left = n.left.insert(t, data)

	case data.Key > n.data.Key:
		n.right = n.right.insert(t, data)

	default:
		return n.rebalance()
	}

	n.level = max(n.left.height(), n.right.height()) + 1
	return n.rebalance()
}

// Η find διατρέχει το δέντρο, αναζητώντας το συγκεκριμένο κλειδί.
func (n *node) find(key int) (Data, error) {
	if n == nil {
		return Data{}, errors.New("key not found")
	}

	switch {
	case n.data.Key == key:
		return n.data, nil

	case key < n.data.Key:
		return n.left.find(key)

	default:
		return n.right.find(key)
	}
}

// Η balRatio παρέχει πληροφορία, σχετικά με τον λόγο ισορροπίας
// του στοιχείου.
func (n *node) balRatio() int {
	return n.right.height() - n.left.height()
}

// Η rotateLeft στρέφει το στοιχείο αριστερά.
//
//	#3          #4
//	  \        /  \
//	  #4      #3  #5
//	    \
//	    #5
func (n *node) rotateLeft() *node {
	r := n.right
	n.right = r.left
	r.left = n
	n.level = max(n.left.height(), n.right.height()) + 1
	r.level = max(r.left.height(), r.right.height()) + 1
	return r
}

// Η rotateRight στρέφει το στοιχείο δεξιά.
//
//	    #5      #4
//	   /       /  \
//	  #4      #3  #5
//	 /
//	#3
func (n *node) rotateRight() *node {
	l := n.left
	n.left = l.right
	l.right = n
	n.level = max(n.left.height(), n.right.height()) + 1
	l.level = max(l.left.height(), l.right.height()) + 1
	return l
}

// Η rotateLeftRight στρέφει το στοιχείο αριστερά και μετά δεξιά.
//
//	  #5          #5      #4
//	 /           /       /  \
//	#3          #4      #3  #5
//	  \        /
//	  #4      #3
func (n *node) rotateLeftRight() *node {
	n.left = n.left.rotateLeft()
	n = n.rotateRight()
	n.level = max(n.left.height(), n.right.height()) + 1
	return n
}

// Η rotateLeftRight στρέφει το στοιχείο αριστερά και μετά δεξιά.
//
//	#3        #3          #4
//	  \         \        /  \
//	  #5        #4      #3  #5
//	 /            \
//	#4            #5
func (n *node) rotateRightLeft() *node {
	n.right = n.right.rotateRight()
	n = n.rotateLeft()
	n.level = max(n.left.height(), n.right.height()) + 1
	return n
}

// Η rebalance περιστρέφει τα στοιχεία, με βάση τον λόγο.
func (n *node) rebalance() *node {
	switch {
	case n.balRatio() < -1 && n.left.balRatio() == -1:
		return n.rotateRight()

	case n.balRatio() > 1 && n.right.balRatio() == 1:
		return n.rotateLeft()

	case n.balRatio() < -1 && n.left.balRatio() == 1:
		return n.rotateLeftRight()

	case n.balRatio() > 1 && n.right.balRatio() == -1:
		return n.rotateRightLeft()
	}
	return n
}

// Η findMax βρίσκει το μέγιστο στοιχείο ενός υπο-δέντρου.
// Η τιμή του αντικαθιστά την τιμή του στοιχείου,
// που θα διαγραφεί. Τιμές επιστροφής: το ίδιο το στοιχείο
// και το πατρικό στοιχείο του.
func (n *node) findMax(parent *node) (*node, *node) {
	switch {
	case n == nil:
		return nil, parent

	case n.right == nil:
		return n, parent
	}
	return n.right.findMax(n)
}

// Η replaceNode αντικαθιστά τον πατρικό δείκτη διεύθυνσης
// του n, με ένα δείκτη διεύθυνσης στο στοιχείο προς αντικατάσταση.
// Το πατρικό στοιχείο δεn πρεπει να έχει τιμή nil.
func (n *node) replaceNode(parent, replacement *node) error {
	if n == nil {
		return errors.New("replaceNode() not allowed on a nil node")
	}

	switch n {
	case parent.left:
		parent.left = replacement

	default:
		parent.right = replacement
	}

	return nil
}

// Η delete απομακρύνει ένα στοιχείο από το δέντρο. Είναι σφάλμα,
// να προσπαθήσει κανείς να διαγραψει ένα στοιχείο, που δεν υπάρχει.
// Προκειμένου να απομακρυνθεί ένα στοιχείο με κατάλληλο τρόπο, η
// Delete χρειάζεται να γνωρίζει το πατρικό στοιχείο του. Αυτο το
// πατρικό στοιχείο, δεν πρέπει να έχει τιμή nil.
func (n *node) delete(key int, parent *node) error {
	if n == nil {
		return errors.New("value to be deleted does not exist in the tree")
	}

	switch {
	case key < n.data.Key:
		return n.left.delete(key, n)

	case key > n.data.Key:
		return n.right.delete(key, n)

	default:
		switch {
		case n.left == nil && n.right == nil:
			n.replaceNode(parent, nil)
			return nil
		case n.left == nil:
			n.replaceNode(parent, n.right)
			return nil
		case n.right == nil:
			n.replaceNode(parent, n.left)
			return nil
		}
		replacement, replParent := n.left.findMax(n)
		n.data = replacement.data
		return replacement.delete(replacement.data.Key, replParent)
	}
}

// Η preOrder διατρέχει ένα στοιχείο διατρέχοντας τα παιδιά του
// με αναδρομή.
func (n *node) preOrder(f func(*node)) {
	if n != nil {
		f(n)
		n.left.preOrder(f)
		n.right.preOrder(f)
	}
}

// Η inOrder διασχίζει το στοιχείο, από το πιο αριστερό στο πιο
// δεξί στοιχείο, ανεξαρτήτως του βάθους.
func (n *node) inOrder(f func(*node)) {
	if n != nil {
		n.left.inOrder(f)
		f(n)
		n.right.inOrder(f)
	}
}

// Η postOrder διασχίζει το στοιχείο από το πιο αριστερό στοιχείο,
// μετά από το αδελφό στοιχείο στο πατρικό στοιχείο, με αναδρομή.
func (n *node) postOrder(f func(*node)) {
	if n != nil {
		n.left.postOrder(f)
		n.right.postOrder(f)
		f(n)
	}
}

// =============================================================================

// Η max επιστρέφει την μεγαλύτερη από τις δύο τιμές.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// =============================================================================

// Η PrettyPrint παίρνει μια τιμή δέντρου και παρουσιάζει
// μια ευχάριστη εκτυπωμένη εκδοχή του δέντρου.
func PrettyPrint(t Tree) {

	// Δημιουργήστε έναν σχεσιακό πίνακα ως ευρετήριο
	// θέσεων για την διάταξη της εκτύπωσης.
	values := make(map[int]int)
	maxIdx := buildIndexMap(values, 0, 0, t.root)

	// Υπολογίστε το συνολικό αριθμό των επιπέδων, με βάση
	// τον μέγιστο παρεχόμενο δείκτη.
	var levels int
	for {
		pow := math.Pow(2, float64(levels))
		if maxIdx < int(pow) {
			break
		}
		levels++
	}
	levels--

	// Κρατήστε τα δεδομένα θέσης προς χρήση.
	data := generateData(levels)

	// Ορίστε την πλευρά της κορυφή του δέντρου.
	for sp := 0; sp < data[0].edge; sp++ {
		fmt.Print(" ")
	}
	fmt.Printf("%02d", values[0])
	fmt.Print("\n")

	dataIdx := 1
	for i := 1; i < len(data); i = i + 2 {

		// Ορίστε την πλευρά αυτής της γραμμής
		for sp := 0; sp < data[i].edge; sp++ {
			fmt.Print(" ")
		}

		// Σχεδιάστε τις γραμμές για αυτή την γραμμή.
		dataHashIdx := dataIdx
		for h := 0; h < data[i].draw; h++ {
			if values[dataHashIdx] != maxInt {
				fmt.Printf("/")
			} else {
				fmt.Printf(" ")
			}
			for sp := 0; sp < data[i].padding; sp++ {
				fmt.Print(" ")
			}
			if values[dataHashIdx+1] != maxInt {
				fmt.Printf("\\")
			} else {
				fmt.Printf(" ")
			}
			dataHashIdx += 2

			if data[i].gaps != 0 && data[i].gaps > h {
				for sp := 0; sp < data[i].gapPad; sp++ {
					fmt.Print(" ")
				}
			}
		}
		fmt.Print("\n")

		// Θέστε την πλευρά της επόμενης πλευράς.
		for sp := 0; sp < data[i+1].edge; sp++ {
			fmt.Print(" ")
		}

		// Σχεδιάστε τους αριθμούς για αυτή την σειρά.
		for n := 0; n < data[i+1].draw; n++ {
			if values[dataIdx] != maxInt {
				fmt.Printf("%02d", values[dataIdx])
			} else {
				fmt.Printf("  ")
			}
			for sp := 0; sp < data[i+1].padding; sp++ {
				fmt.Print(" ")
			}
			if values[dataIdx+1] != maxInt {
				fmt.Printf("%02d", values[dataIdx+1])
			} else {
				fmt.Printf("  ")
			}
			dataIdx += 2

			if data[i+1].gaps != 0 && data[i+1].gaps > n {
				for sp := 0; sp < data[i+1].gapPad; sp++ {
					fmt.Print(" ")
				}
			}
		}
		fmt.Print("\n")
	}

	fmt.Print("\n")
}

const maxInt = int(^uint(0) >> 1)

// Η buildIndex διατρέχει το δέντρο και παράγει έναν σχεσιακό
// πίνακα δεικτών θέσεων, για κάθε στοιχείο στο
// δέντρο, προς εκτύπωση.
//
//	      40
//	   /      \
//	  05      80
//	 /  \    /  \
//	02  25  65  98
//
// values{0:40, 1:05, 2:80, 3:02, 4:25, 5:65, 6:98}
func buildIndexMap(values map[int]int, idx int, maxIdx int, n *node) int {

	// Πρέπει να κρατάμε λογαριασμό της υψηλότερης θέσης δείκτη
	// που χρησιμοποιείται, προκειμένου να υποβοηθηθεί ο υπολογισμός
	// του βάθους του δέντρου.
	if idx > maxIdx {
		maxIdx = idx
	}

	// Φτάσαμε στο τέλος ενός τμήματος. Χρησιμοποιήστε την maxInt προκειμένου
	// να σηματοδοτηθει, ότι δεν υπάρχει τιμή σε αυτή την θέση.
	if n == nil {
		values[idx] = maxInt
		return maxIdx
	}

	// Αποθηκεύστε την τιμή αυτού του στοιχείου στον σχεσιακό πίνακα,
	// στην υπολογισμένη θέση δείκτη.
	values[idx] = n.data.Key

	// Ελεγξτε αν υπάρχουν ακόμα στοιχεία προς έλεγχο στο αριστερό
	// τμήμα. Όταν μετακινούμαστε προς τα κάτω στο δέντρο, διπλασιάζεται
	// ο επόμενος δείκτης.
	if n.left != nil {
		nextidx := 2*idx + 1
		maxIdx = buildIndexMap(values, nextidx, maxIdx, n.left)
	}

	// Ελέγξτε αν υπάρχουν ακόμα στοιχεία προς έλεγχο, προς τα κάτω
	// στο δεξί τμήμα. Όταν μετακινούμαστε προς τα κάτω στο δέντρο, διπλασιάζεται
	// ο επόμενος δείκτης.
	nextidx := 2*idx + 2
	maxIdx = buildIndexMap(values, nextidx, maxIdx, n.right)

	// Πρέπει να θέσουμε τους απόντες δείκτες στον σχεσιακό πίνακα
	// στην τιμή maxInt. Με αυτό τον τρόπο, θα αγνοηθούν κατά την εκτύπωση.
	if idx == 0 {
		for i := 0; i < maxIdx; i++ {
			if _, ok := values[i]; !ok {
				values[i] = maxInt
			}
		}
	}

	return maxIdx
}

// Ο pos παρέχει δεδομένα θέσης, για την εκτύπωση ενός δέντρου.
type pos struct {
	edge    int
	draw    int
	padding int
	gaps    int
	gapPad  int
}

// Η generateData παράγει όλα τα δεδομένα θέσης που χρειάζονται,
// προκειμένου να παρουσιαστούν τα στοιχεία, σε διαφορετικά επίπεδα.
func generateData(level int) []pos {
	totalData := (level * 2) - 1
	data := make([]pos, totalData)
	edge := 1
	draw := level - 2
	padding := 0
	gapPad := 2

	for i := totalData - 1; i >= 0; i = i - 2 {

		// Δημιουργήστε τις αρχικές θέσεις των πλευρών.
		data[i].edge = int(math.Pow(2, float64(edge)))
		if i > 0 {
			data[i-1].edge = data[i].edge + 1
		}
		edge++

		// Δημιουργήστε πληροφορίες σχεδιασμού.
		if draw > 0 {
			data[i].draw = int(math.Pow(2, float64(draw)))
			data[i-1].draw = data[i].draw
		} else {
			data[i].draw = 1
			if i > 0 {
				data[i-1].draw = 1
			}
		}
		draw--

		// Δημιουργήστε πληροφορίες ενθεμάτων.
		padding += data[i].edge
		data[i].padding = padding
		if i > 0 {
			data[i-1].padding = padding
		}

		// Δημιουργήστε πληροφορίες κενών.
		data[i].gaps = data[i].draw - 1
		if i > 0 {
			data[i-1].gaps = data[i].gaps
		}

		// Δημιουργήστε πληροφορίες ενθεμάτων κενών.
		if i > 2 {
			data[i-1].gapPad = int(math.Pow(2, float64(gapPad)))
			data[i].gapPad = data[i-1].gapPad - 2
		}
		gapPad++
	}

	return data
}
