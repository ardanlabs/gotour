//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Αλλάξτε τον κώδικα από το πρότυπο. Προσθέστε ένα νέο τύπο CachingFeed που ενσωματώνει τον Feed και
// παρακάμπτει την μέθοδο τύπου Fetch.

// Αυτό το πρόγραμμα ορίζει έναν τύπο Feed με δύο μεθόδους: την Count και την Fetch. Δημιουργείστε έναν
// νέο τύπο CachingFeed που ενσωματώνει τον *Feed αλλά ξανά ορίζει (override) την μέθοδο τύπου Fetch.
//
// Ο τύπος CachingFeed θα πρέπει να έχει να έχει έναν πίνακα κατακερματισμού με Documents προκειμένου να
// περιοριστεί το πλήθος των κλήσεων στην Feed.Fetch.
package main

import (
	"fmt"
	"log"
	"time"
)

// Ο Document είναι το βασικό μοντέλο των δεδομένων με το οποίο δουλεύουμε.
type Document struct {
	Key   string
	Title string
}

// ==================================================

// Ο Feed είναι ένας τύπος που γνωρίζει πως να παραλαμβάνει (fetch) Document.
type Feed struct{}

// H Count ενημερώνει πόσα έγγραφα είναι διαθέσιμα προς παραλαβή.
func (f *Feed) Count() int {
	return 42
}

// Η Fetch μιμείται την αναζήτηση ενός εγγράφου με βάση το κλειδί. Είναι αργή.
func (f *Feed) Fetch(key string) (Document, error) {
	time.Sleep(time.Second)

	doc := Document{
		Key:   key,
		Title: "Title for " + key,
	}
	return doc, nil
}

// ==================================================

// Ο FetchCounter είναι η συμπεριφορά στην οποία εξαρτόμαστε για την συνάρτηση επεξεργασίας μας.
type FetchCounter interface {
	Fetch(key string) (Document, error)
	Count() int
}

func process(fc FetchCounter) {
	fmt.Printf("There are %d documents\n", fc.Count())

	keys := []string{"a", "a", "a", "b", "b", "b"}

	for _, key := range keys {
		doc, err := fc.Fetch(key)
		if err != nil {
			log.Printf("Could not fetch %s : %v", key, err)
			return
		}

		fmt.Printf("%s : %v\n", key, doc)
	}
}

// ==================================================

// Ο CachingFeed κρατάει ένα τοπικό αντίγραφο των Document που έχουν ήδη
// ανασυρθεί. Ενσωματώνει τον Feed προκειμένου να αποκτήσει την συμπεριφορά Fetch και την συμπεριφορά Count
// όμως ξανά ορίζει την Fetch προκειμένου να έχει την μνήμη αποθήκευσης.
type CachingFeed struct {
	// TODO ενσωματώστε τον *Feed και προσθέστε ένα πεδίο για έναν map[string]Document.
}

// Η NewCachingFeed δίνει αρχική τιμή σε έναν CachingFeed έτοιμο προς χρήση.
func NewCachingFeed(f *Feed) *CachingFeed {

	// TODO δημιουργείστε έναν CachingFeed με έναν πίνακα κατακερματισμού με αρχική τιμή και
	// ενσωματωμένο feed.
	// Επιστρέψτε την διεύθυνση του.

	return nil // Απομακρύνετε αυτή την γραμμή.

}

// Η Fetch καλεί την μέθοδο τύπου Fetch του ενσωματωμένου τύπου αν το κλειδί δεν είναι στην προσωρινή μνήμη.
func (cf *CachingFeed) Fetch(key string) (Document, error) {

	// TODO υλοποιείστε αυτή την μέθοδο τύπου. Ελέγξτε το πεδίο του πίνακα
	// κατακερματισμού για το συγκεκριμένο κλειδί και επιστρέψτε το αν βρεθεί.
	// Αν δεν βρεθεί, καλέστε την μέθοδο τύπου Fetch του ενσωματωμένου τύπου.
	// Αποθηκεύστε το αποτέλεσμα στον πίνακα κατακερματισμού πριν το επιστρέψετε.

	return Document{}, nil // Απομακρύνετε αυτή την γραμμή.
}

// ==================================================

func main() {
	fmt.Println("Using Feed directly")
	process(&Feed{})

	// Καλέστε την process ξανά με τον δικό σας CachingFeed.
	//fmt.Println("Using CachingFeed")
}
