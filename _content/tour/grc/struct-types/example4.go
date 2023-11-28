//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://github.com/dominikh/go-tools#installation
// go install honnef.co/go/tools/cmd/...@2023.1.3

// Η Ευθυγράμμιση/στοίχιση αφορά την τοποθέτηση πεδίων εντός ευθυγραμμισμένων ορίων
// μνήμης, για πιο αποτελεσματικές αναγνώσεις και εγγραφές στην μνήμη.

// Δείγμα προγράμματος προκειμένου να παρουσιαστεί ο τρόπος με τον οποίο, οι τύποι
// ευθυγραμμίζονται κατά μήκος ορίων μνήμης.
package main

import (
	"fmt"
	"unsafe"
)

// Χωρίς προσθήκη byte.
type nbp struct {
	a bool // 	1 byte				μέγεθος 1
	b bool // 	1 byte				μέγεθος 2
	c bool // 	1 byte				μέγεθος 3 - Ευθυγράμμιση σε 1 byte
}

// Προσθήκη ενός byte.
type sbp struct {
	a bool //	1 byte				μέγεθος 1
	//			προσθήκη 1 byte		μέγεθος 2
	b int16 // 	2 byte				μέγεθος 4 - Ευθυγράμμιση σε 2 byte
}

// Προσθήκη τριών byte.
type tbp struct {
	a bool //	1 byte				μέγεθος 1
	//			προσθήκη 3 byte		μέγεθος 4
	b int32 //	4 byte				μέγεθος 8 - Ευθυγράμμιση σε 4 byte
}

// Προσθήκη εφτά byte.
type svnbp struct {
	a bool //	1 byte				μέγεθος 1
	//			προσθήκη 7 byte		μέγεθος 8
	b int64 //	8 byte				μέγεθος 16 - Ευθυγράμμιση σε 8 byte
}

// Χωρίς προσθήκη.
type np struct {
	a string // 16 byte			μέγεθος 16
	b string // 16 byte			μέγεθος 32
	c int32  //  4 byte			μέγεθος 36
	d int32  //  4 byte			μέγεθος 40 - Ευθυγράμμιση σε 8 byte
}

// Προσθήκη οκτώ byte σε αρχιτεκτονική υλικού 64bit. Το μέγεθος λέξης είναι 8 byte.
type ebp64 struct {
	a string //	16 byte				μέγεθος 16
	b int32  //	 4 byte				μέγεθος 20
	//  		 4 byte προσθήκη	μέγεθος 24
	c string //	16 byte				μέγεθος 40
	d int32  //	 4 byte		 		μέγεθος 44
	//  		 4 byte προσθήκη	μέγεθος 48 - Ευθυγράμμιση σε 8 byte
}

func main() {
	var nbp nbp
	size := unsafe.Sizeof(nbp)
	fmt.Printf("nbp  : SizeOf[%d][%p %p %p]\n", size, &nbp.a, &nbp.b, &nbp.c)

	// -------------------------------------------------------------------------

	var sbp sbp
	size = unsafe.Sizeof(sbp)
	fmt.Printf("sbp  : SizeOf[%d][%p %p]\n", size, &sbp.a, &sbp.b)

	// -------------------------------------------------------------------------

	var tbp tbp
	size = unsafe.Sizeof(tbp)
	fmt.Printf("tbp  : SizeOf[%d][%p %p]\n", size, &tbp.a, &tbp.b)

	// -------------------------------------------------------------------------

	var svnbp svnbp
	size = unsafe.Sizeof(svnbp)
	fmt.Printf("svnbp: SizeOf[%d][%p %p]\n", size, &svnbp.a, &svnbp.b)

	// -------------------------------------------------------------------------

	var np np
	size = unsafe.Sizeof(np)
	fmt.Printf("np   : SizeOf[%d][%p %p %p %p]\n", size, &np.a, &np.b, &np.c, &np.d)

	// -------------------------------------------------------------------------

	var ebp64 ebp64
	size = unsafe.Sizeof(ebp64)
	fmt.Printf("ebp64: SizeOf[%d][%p %p %p %p]\n", size, &ebp64.a, &ebp64.b, &ebp64.c, &ebp64.d)
}
