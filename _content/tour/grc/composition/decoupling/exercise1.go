//go:build OMIT || nobuild

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Χρησιμοποιώντας το πρότυπο, δηλώστε ένα σύνολο πραγματικών τύπων, που υλοποιούν
// ένα σύνολο προκαθορισμένων τύπων διεπαφών. Στην συνέχεια, δημιουργείστε τιμές 
// αυτών των τύπων και χρησιμοποιήστε τις, προκειμένου να ολοκληρώσετε ένα σύνολο
// από προκαθορισμένες ενέργειες.
package main

// Προσθέστε δηλώσεις εισαγωγής (import).

// Ο administrator αναπαριστά ένα πρόσωπο ή άλλη οντότητα, που είναι ικανοί
// να διαχειρίζονται υποδομές υλικού και λογισμικού.
type administrator interface {
	administrate(system string)
}

// Ο developer αναπαριστά ένα πρόσωπο ή άλλη οντότητα, που μπορούν να γράφουν
// λογισμικό.
type developer interface {
	develop(system string)
}

// =============================================================================

// Ο adminlist αναπαριστά ένα σύνολο από administrator.
type adminlist struct {
	list []administrator
}

// Η Enqueue προσθέτει έναμ administrator σε έναν adminlist.
func (l *adminlist) Enqueue(a administrator) {
	l.list = append(l.list, a)
}

// Η Dequeue απομακρύνει έναν administrator από έναν adminlist.
func (l *adminlist) Dequeue() administrator {
	a := l.list[0]
	l.list = l.list[1:]
	return a
}

// =============================================================================

// Ο devlist αναπαριστά ένα σύνολο από developer.
type devlist struct {
	list []developer
}

// Η Enqueue προσθέτει έναν developer στον devlist.
func (l *devlist) Enqueue(d developer) {
	l.list = append(l.list, d)
}

// Η Dequeue απομακρύνει έναν developer από τον devlist.
func (l *devlist) Dequeue() developer {
	d := l.list[0]
	l.list = l.list[1:]
	return d
}

// =============================================================================

// Δηλώστε έναν πραγματικό τύπο με το όνομα sysadmin, με ένα πεδίο name 
// τύπου συμβολοσειράς.

// Δηλώστε μια μέθοδο τύπου με το όνομα administrate για τον τύπο sysadmintype, 
// που υλοποιεί την διεπαφή administrator. Η administrate θα πρέπει να τυπώνει 
// το πεδίο name του sysadmin, όπως επίσης και το σύστημα το οποίο διαχειρίζεται.

// Δηλώστε έναν πραγματικό τύπο με το όνομα programmer, με ένα πεδίο name τύπου
// συμβολοσειράς.

// Δηλώστε μια μέθοδο με το όνομα develop για τον τύπο programmer, υλοποιώντας 
// την διεπαφή developer. Η develop πρέπει να τυπώνει το πεδίο name του 
// programmer, όπως επίσης το σύστημα που γράφουν.

// Δηλώστε έναν πραγματικό τύπο με το όνομα company. Δηλώστε τον σαν την σύνθεση 
// των τύπων διεπαφής administrator και developer.

// =============================================================================

func main() {

	// Δημιουργείστε μια μεταβλητή με το όνομα admins τύπου adminlist.

	// Δημιουργείστε μια μεταβλητή με το όνομα devs τύπου devlist.

	// Με την Enqueue προσθέστε έναν νέο sysadmin στον admins.

	// Με την Enqueue προσθέστε δύο νέους programmers στον devs.

	// Δημιουργείστε μια μεταβλητή με το όνομα cmp τύπου company, και δώστε 
	// αρχική τιμή προσλαμβάνονατς (στμ. dequeuing) έναν administrator από τον admins και
	// έναν developer από τον devs.

	// Με την Enqueue προσθέστε την τιμή company και στις δύο λίστες καθώς ο company 
	// υλοποιεί κάθε διεπαφή.

	// Ένα σύνολο εργασιών προς ολοκλήρωση από administrator και developer.
	tasks := []struct {
		needsAdmin bool
		system     string
	}{
		{needsAdmin: false, system: "xenia"},
		{needsAdmin: true, system: "pillar"},
		{needsAdmin: false, system: "omega"},
	}

	// Προσπελάστε επαναλληπτικά τις εργασίες.
	for _, task := range tasks {

		// Ελέγξτε αν η εργασία χρειάζεται έναν administrator, διαφορετικά 
		// χρησιμοποιήστε έναν developer.
		if {

			// Με την Dequeue απομακρύνετε μια τιμή administrator από 
			// την λίστα admins και καλέστε την μέθοδο τύπου administrate.

			continue
		}

		// Με την Dequeue απομακρύνετε μια τιμή developer από την λίστα devs και 
		// καλέστε την μέθοδο τύπου develop.
	}
}
