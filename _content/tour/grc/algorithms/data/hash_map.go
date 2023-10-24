//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Αυτό το δείγμα προγράμματος παρουσιάζει τον τρόπο συγγραφής ενός βασικού
// πίνακα κατακερματισμού.
package main

import (
	"fmt"
	"hash/maphash"
)

func main() {
	h := New()

	k1, v1 := "key1", 1
	k2, v2 := "key2", 2
	h.Store(k1, v1)
	h.Store(k2, v2)

	v, err := h.Retrieve(k1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("key:", k1, "value:", v)

	v1b := 11
	h.Store(k1, v1b)

	v, err = h.Retrieve(k1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("key:", k1, "value:", v)

	if err := h.Delete(k1); err != nil {
		fmt.Println(err)
		return
	}

	v, err = h.Retrieve(k1)
	if err != nil {
		fmt.Println(err)
	}

	fn := func(key string, value int) bool {
		fmt.Println("key:", key, "value:", value)
		return true
	}
	h.Do(fn)
}

// =============================================================================

const numBuckets = 256

// Ένας entry όπου αποθηκεύουμε κλειδί και τιμή στον πίνακα.
type entry struct {
	key   string
	value int
}

// Ο Hash είναι μια απλή υλοποίηση πίνακα κατακερματισμού.
type Hash struct {
	buckets [][]entry
	hash    maphash.Hash
}

// Η New επιστρέφει έναν νέο πίνακα κατακερματισμού.
func New() *Hash {
	return &Hash{
		buckets: make([][]entry, numBuckets),
	}
}

// Η Store προσθέτει μια τιμή στον πίνακα κατακερματισμού,
// με βάση το κλειδί.
func (h *Hash) Store(key string, value int) {

	// Για το συγκεκριμένο κλειδί, προσδιορίστε σε ποια
	// θέση αποθήκευσης στην φέτα πρέπει να αποθηκεύσουμε
	// τα κλειδί/τιμή.
	idx := h.hashKey(key)

	// Εξάγετε ένα αντίγραφο από την θέση αποθήκευσης
	// από τον πίνακα κατακερματισμού.
	bucket := h.buckets[idx]

	// Επαναλάβετε επί των δεικτών για την συγκεκριμένη
	// θέση αποθήκευσης.
	for idx := range bucket {

		// Συγκρίνετε τα κλειδιά και αν υπάρχει κάποιο που ταιριάζει,
		// αντικαταστήστε την υπάρχουσα τιμή της εγγραφής με την νέα.
		if bucket[idx].key == key {
			bucket[idx].value = value
			return
		}
	}

	// Αυτό το κλειδί δεν υπάρχει, επομένως προσθέστε αυτή την νέα τιμή.
	h.buckets[idx] = append(bucket, entry{key, value})
}

// Η Retrieve εξάγει μια τιμή από τον πίνακα κατακερματισμού
// με βάση το κλειδί.
func (h *Hash) Retrieve(key string) (int, error) {

	// Για το συγκεκριμένο κλειδί, προσδιορίστε σε ποια θέση
	// αποθήκευσης μέσα στην φέτα πρέπει να αποθηκεύσουμε
	// τα κλειδί/τιμή.
	idx := h.hashKey(key)

	// Επαναλάβετε επί των εγγραφών για την συγκεκριμένη θέση
	// αποθήκευσης.
	for _, entry := range h.buckets[idx] {

		// Συγκρίνετε τα κλειδιά και αν υπάρχει ταύτιση,
		// επιστρέψτε την τιμή που σχετίζεται με το κλειδί.
		if entry.key == key {
			return entry.value, nil
		}
	}

	// Το κλειδί δεν βρέθηκε, επομένως επιστρέψτε το σφάλμα.
	return 0, fmt.Errorf("%q not found", key)
}

// Η Delete διαγράφει μια εγγραφή, από τον πίνακα κατακερματισμού.
func (h *Hash) Delete(key string) error {

	// Για το συγκεκριμένο κλειδί, προσδιορίστε σε ποια θέση
	// αποθήκευσης στην φέτα πρέπει να αποθηκεύσουμε τα κλειδί/τιμή.
	bucketIdx := h.hashKey(key)

	// Εξάγετε ένα αντίγραφο της θέσης αποθήκευσης, από τον πίνακα
	// κατακερματισμού.
	bucket := h.buckets[bucketIdx]

	// Επαναλάβετε επί των εγγραφών για την συγκεκριμένη
	// θέση αποθήκευσης.
	for entryIdx, entry := range bucket {

		// Συγκρίνετε τα κλειδιά και αν υπάρχει ταύτιση, απομακρύνετε
		// την εγγραφή από την θέση αποθήκευσης.
		if entry.key == key {

			// Απομακρύνετε την εγγραφή με βάση την θέση δείκτη της.
			bucket = removeEntry(bucket, entryIdx)

			// Αντικαταστήστε τον υπάρχοντα χώρο αποθήκευσης
			h.buckets[bucketIdx] = bucket
			return nil
		}
	}

	// Το κλειδί δεν βρέθηκε, επομένως επιστρέψτε το σφάλμα.
	return fmt.Errorf("%q not found", key)
}

// Η Len επιστρέφει τον αριθμό των στοιχείων στον πίνακα κατακερματισμού.
// Η συνάρτηση, στην τρέχουσα μορφή της, χρησιμοποιεί μια γραμμική
// προσπέλαση η οποία όμως θα μπορούσε να βελτιωθεί, με μετα-δεδομένα.
func (h *Hash) Len() int {
	sum := 0
	for _, bucket := range h.buckets {
		sum += len(bucket)
	}
	return sum
}

// Η Do καλεί την fn σε κάθε κλειδί/τιμή. Αν η fn επιστρέψει false,
// τότε σταματάει η επαναληπτική προσπέλαση.
func (h *Hash) Do(fn func(key string, value int) bool) {
	for _, bucket := range h.buckets {
		for _, entry := range bucket {
			if ok := fn(entry.key, entry.value); !ok {
				return
			}
		}
	}
}

// Η hashKey υπολογίζει ποια θέση δείκτη της θέσης αποθήκευσης
// να χρησιμοποιήσει, για το συγκεκριμένο κλειδί.
func (h *Hash) hashKey(key string) int {

	// Καλέστε την Reset στον maphash του πίνακα
	// στην αρχική κατάσταση, ώστε να πάρουμε την ίδια
	// τιμή κατακερματισμού για το ίδιο κλειδί.
	h.hash.Reset()

	// Γράψτε το κλειδί στον maphash, προκειμένου να ανανεωθεί
	// η παρούσα κατάσταση. Δε ελέγχουμε την τιμή σφάλματος
	// καθώς η WriteString δεν αποτυγχάνει ποτε.
	h.hash.WriteString(key)

	// Αναζητήστε από στον maphash την τρέχουσα κατάσταση
	// την οποία θα χρησιμοποιήσουμε, προκειμένου να υπολογίσουμε
	// τον τελικό δείκτη της θέσης αποθήκευσης.
	n := h.hash.Sum64()

	// Χρησιμοποιήστε τον τελεστή ακέραιας διαίρεσης (στμ. modulus)
	// προκειμένου να επιστρέψετε μια τιμή στο εύρος
	// του διαθέσιμου μήκους για τις θέσεις αποθήκευσης,
	// όπως ορίζεται από την σταθερά numBuckets.
	return int(n % numBuckets)
}

// Η removeEntry πραγματοποιεί την απομάκρυνση μιας εγγραφής
// από μια θέση αποθήκευσης.
func removeEntry(bucket []entry, idx int) []entry {

	// https://github.com/golang/go/wiki/SliceTricks
	// Απομακρύνετε την εγγραφή παίρνοντας όλες τις
	// εγγραφές μπροστά από τον δείκτη θέσης και
	// μεταφέροντας τες πίσω από τον συγκεκριμένο δείκτη.
	copy(bucket[idx:], bucket[idx+1:])

	// Θέστε το κατάλληλο μήκος για την νέα φέτα καθώς
	// απομακρύνθηκε μια εγγραφή. Το μήκος πρέπει να
	// μειωθεί κατά 1.
	bucket = bucket[:len(bucket)-1]

	// Προσπαθήστε να διαπιστώσετε αν η τρέχουσα εκχώρηση
	// για την θέση αποθήκευσης μπορεί να μειωθεί,
	// εξαιτίας του αριθμού των εγγραφών που αφαιρέθηκαν
	// από την θέση αποθήκευσης.
	return reduceAllocation(bucket)
}

// Η reduceAllocation αναζητά αν μπορεί να απελευθερώσει
// μνήμη, όταν ένας χώρος αποθήκευσης έχει χάσει κάποιο
// ποσοστό των εγγραφών του.
func reduceAllocation(bucket []entry) []entry {

	// Αν η θέση αποθήκευσης είναι περισσότερο γεμάτη από το ½,
	// τότε μην κάνετε τίποτα.
	if cap(bucket) < 2*len(bucket) {
		return bucket
	}

	// Απελευθερώστε την μνήμη όταν ο χώρος αποθήκευσης συρρικνώνεται
	// πολύ. Αν δεν το κάνουμε αυτό, ο υποκείμενος πίνακας αποθήκευσης
	// θα παραμείνει στην μνήμη και θα συνεχίσει να έχει το μεγαλύτερο
	// μέγεθος, που είχε ποτέ η θέση αποθήκευσης.
	newBucket := make([]entry, len(bucket))
	copy(newBucket, bucket)
	return newBucket
}
