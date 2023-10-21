//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος προκειμένου να παρουσιαστεί ο τρόπος αποθήκευσης και ανάσυρσης
// τιμών από ένα context.
package main

import (
	"context"
	"fmt"
)

// Ο TraceID αναπαριστά ένα αναγνωριστικό ίχνους.
type TraceID string

// Ο TraceIDKey είναι ο τύπος τιμής προς χρήση για το κλειδί. Το κλειδί είναι
// συγκεκριμένου τύπου και μόνο τιμές του ίδιου τύπου θα ταιριάζουν.
type TraceIDKey int

func main() {

	// Δημιουργείστε μια traceID για αυτό το αίτημα.
	traceID := TraceID("f47ac10b-58cc-0372-8567-0e02b2c3d479")

	// Δηλώστε ένα κλειδί με την τιμή μηδέν τύπου userKey.
	const traceIDKey TraceIDKey = 0

	// Αποθηκεύστε την τιμή traceID στην context με τιμή μηδέν
	// για τον τύπο κλειδιού.
	ctx := context.WithValue(context.Background(), traceIDKey, traceID)

	// Ανασύρετε αυτή την τιμή traceID από την θήκη τιμών της Context.
	if uuid, ok := ctx.Value(traceIDKey).(TraceID); ok {
		fmt.Println("TraceID:", uuid)
	}

	// Ανασύρετε αυτή την τιμή traceID από την θήκη τιμών της Context χωρίς
	// να χρησιμοποιείτε τον κατάλληλο τύπο κλειδιού.
	if _, ok := ctx.Value(0).(TraceID); !ok {
		fmt.Println("TraceID Not Found")
	}
}
