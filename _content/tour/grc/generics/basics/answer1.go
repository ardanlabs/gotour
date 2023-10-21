//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Υλοποιεί μια συνάρτηση γενικού προγραμματισμού που μπορεί να σειριοποιήσει (marshal) σε JSON.
package main

import (
	"encoding/json"
	"fmt"
)

// Υλοποιείστε την συνάρτηση γενικού προγραμματισμού με το όνομα
// marshal που μπορεί να δεχθεί κάθε τιμή τύπου T και να σειριοποιήσει
// αυτή την τιμή σε JSON.
func marshal[T any](v T) ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Δηλώστε έναν τύπο struct με το όνομα User με δύο πεδία, το πεδίο Name και το πεδίο Age.
type User struct {
	Name string
	Age  int
}

func main() {

	// Δημιουργήστε μια τιμή τύπου User.
	u := User{
		Name: "Bill",
		Age:  10,
	}

	// Καλέστε την συνάρτηση γενικού προγραμματισμού marshal.
	data, err := marshal(u)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Εκτυπώστε την JSON που παράγει η συνάρτηση marshal.
	fmt.Println(string(data))
}
