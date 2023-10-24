//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος, προκειμένου να παρουσιαστεί πως μπορεί να χρησιμοποιηθεί
// η ενσωμάτωση τύπων, προκειμένου να επαναχρησιμοποιηθούν συμπεριφορές από
// άλλους τύπους και να ξανάοριστούν συγκεκριμένες μέθοδοι τύπου.
package main

import (
	"fmt"
	"log"
	"time"
)

// Ο Document είναι το βασικό μοντέλο δεδομένων, με το οποίο εργαζόμαστε.
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

// Ο CachingFeed κρατάει ένα τοπικό αντίγραφο των Document, που έχουν ήδη
// ανασυρθεί. Ενσωματώνει τον Feed προκειμένου να αποκτήσει την συμπεριφορά
// Fetch και την συμπεριφορά Count, όμως ξαναορίζει την Fetch, προκειμένου
// να έχει την μνήμη αποθήκευσης.
type CachingFeed struct {
	docs map[string]Document
	*Feed
}

// Η NewCachingFeed δίνει αρχική τιμή σε έναν CachingFeed έτοιμο προς χρήση.
func NewCachingFeed(f *Feed) *CachingFeed {
	return &CachingFeed{
		docs: make(map[string]Document),
		Feed: f,
	}
}

// Η Fetch καλεί την μέθοδο τύπου Fetch του ενσωματωμένου τύπου, αν το κλειδί
// δεν είναι στην προσωρινή μνήμη.
func (cf *CachingFeed) Fetch(key string) (Document, error) {
	if doc, ok := cf.docs[key]; ok {
		return doc, nil
	}

	doc, err := cf.Feed.Fetch(key)
	if err != nil {
		return Document{}, err
	}

	cf.docs[key] = doc
	return doc, nil
}

// ==================================================

// Ο FetchCounter είναι η συμπεριφορά στην οποία εξαρτόμαστε για την συνάρτηση
// επεξεργασίας μας.
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

func main() {
	fmt.Println("Using Feed directly")
	process(&Feed{})

	fmt.Println("Using CachingFeed")
	c := NewCachingFeed(&Feed{})
	process(c)
}
