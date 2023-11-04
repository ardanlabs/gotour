//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Υλοποιήστε μια γενική συνάρτηση με το όνομα marshal, που να μπορεί να
// σειριοποιήσει (στμ. marshal) JSON όμως δέχεται μόνο τιμές, που υλοποιούν
// την διεπαφή json.Marshaler interface.
package main

import (
	"encoding/json"
	"fmt"
)

// Υλοποιήστε μια γενική συνάρτηση με το όνομα marshal, που μπορεί να
// αποδέχεται μόνο τιμές τύπου T, που υλοποιεί την διεπαφή json.Marshaler.
func marshal[T json.Marshaler](v T) ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// =============================================================================

// Ορίστε έναν τύπο με το όνομα user με δύο πεδία, το πεδίο name και
// το πεδίο email.
type user struct {
	name  string
	email string
}

// Υλοποιήστε μια μέθοδο τύπου, που υλοποιεί με την σειρά της την
// διεπαφή json.Marshaler. Η μέθοδος τύπου πρέπει να επιστρέφει μια
// τιμή τύπου user ως JSON.
func (u user) MarshalJSON() ([]byte, error) {
	v := fmt.Sprintf("{\"name\": %q, \"email\": %q}", u.name, u.email)
	return []byte(v), nil
}

// =============================================================================

func main() {

	// Δημιουργείστε μια τιμή τύπου user.
	user := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	// Καλέστε την συνάρτηση γενικού προγραμματισμού marshal.
	s1, err := marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Παρουσιάστε την επιστραμμένη τιμή JSON.
	fmt.Println("user:", string(s1))
}
