//go:build OMIT || norun

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος, που υλοποιεί μια απλή υπηρεσία ιστού (στμ. web service)
// χρησιμοποιώντας το context, προκειμένου να διαχειριστούμε τον χρόνο ακύρωσης
// και περνώντας το context στο αίτημα.
package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Ο τύπος key είναι μη εξαγόμενος, προκειμένου να αποφευχθούν συγκρούσεις
// με κλειδιά στο context, που είναι ορισμένα σε άλλα πακέτα.
type key int

// Η userIPkey είναι το κλειδί του context για την διεύθυνση IP του χρήστη.
// Η τιμή μηδέν είναι αυθαίρετη. Αν αυτό το πακέτο όριζε άλλα κλειδιά του
// context, θα είχαν διαφορετικές ακέραιες τιμές.
const userIPKey key = 0

// Ο User ορίζει έναν χρήστη στο σύστημα.
type User struct {
	Name  string
	Email string
}

func main() {
	routes()

	log.Println("listener : Started : Listening on: http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}

// Η routes ορίζει τις διαδρομές της υπηρεσίας ιστού.
func routes() {
	http.HandleFunc("/user", findUser)
}

// Η findUser κάνει μια κλήση βάσης δεδομένων, προκειμένου να βρεθεί ένας χρήστης.
func findUser(rw http.ResponseWriter, r *http.Request) {

	// Δημιουργήστε ένα context με χρόνο ακύρωσης πενήντα millisecond.
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	// Αποθηκεύστε την διεύθυνσης ip του χρήστη στο context. Αυτή η
	// κλήση επιστρέφει ένα νέο context, το οποίο πρέπει τώρα να
	// χρησιμοποιήσουμε. Το αρχικό context είναι αυτό από το οποίο
	// παράγεται το νέο context.
	ctx = context.WithValue(ctx, userIPKey, r.RemoteAddr)

	// Χρησιμοποιήστε αυτή την ρουτίνα συνεκτέλεσης της Go προκειμένου
	// να πραγματοποιήσετε την κλήση στην βάση δεδομένων. Χρησιμοποιήστε
	// το κανάλι επικοινωνίας, προκειμένου να λάβετε τον χρήστη πίσω.
	ch := make(chan *User, 1)
	go func() {

		// Πάρτε την διεύθυνση ip από το context για καταγραφή.
		if ip, ok := ctx.Value(userIPKey).(string); ok {
			log.Println("Start DB for IP", ip)
		}

		// Πραγματοποιείστε την κλήσης της βάσης δεδομένων και επιστρέψτε
		// την τιμή στο κανάλι επικοινωνίας.
		ch <- readDatabase()
		log.Println("DB goroutine terminated")
	}()

	// Περιμένετε ώστε η κλήση στην βάση δεδομένων να τελειώσει ή να παρέλθει
	// ο διαθέσιμος χρόνος.
	select {
	case u := <-ch:

		// Απαντήστε με τον χρήστη.
		sendResponse(rw, u, http.StatusOK)
		log.Println("Sent StatusOK")
		return

	case <-ctx.Done():

		// Αν έχετε την δυνατότητα να ακυρώσετε την λειτουργία
		// στην βάση δεδομένων, που πραγματοποιεί η ρουτίνα
		// συνεκτέλεσης της Go, πραγματοποιείστε το τώρα.
		// Σε αυτό το παράδειγμα, δεν μπορούμε.

		// Απαντήστε με το σφάλμα.
		e := struct{ Error string }{ctx.Err().Error()}
		sendResponse(rw, e, http.StatusRequestTimeout)
		log.Println("Sent StatusRequestTimeout")
		return
	}
}

// Η readDatabase πραγματοποιεί μια προσποιητή κλήση βάσης δεδομένων,
// με καθυστέρηση ενός second.
func readDatabase() *User {
	u := User{
		Name:  "Bill",
		Email: "bill@ardanlabs.com",
	}

	// Δημιουργήστε καθυστέρηση 100 millisecond.
	time.Sleep(100 * time.Millisecond)

	return &u
}

// Η sendResponse σειριοποιεί (marshal) την παρεχόμενη τιμή σε μορφή json και
// επιστρέφει αυτή την τιμή στον καλώντα.
func sendResponse(rw http.ResponseWriter, v interface{}, statusCode int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	json.NewEncoder(rw).Encode(v)
}
