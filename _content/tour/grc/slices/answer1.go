//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δηλώστε μια nil φέτα ακεραίων. Δημιουργείστε μια επανάληψη, που προσθέτει 10 τιμές
// στην φέτα. Επισκεφθείτε διαδοχικά τα στοιχεία της φέτας και παρουσιάστε κάθε τιμή.
//
// Δηλώστε μια φέτα πέντε συμβολοσειρών και δώστε αρχική τιμή στην φέτα με ρητές τιμές
// συμβολοσειρών. Παρουσιάστε όλα τα στοιχεία. Δημιουργείστε μια φέτα των δεικτών ένα
// και δύο και παρουσιάστε την θέση δείκτη και την τιμή κάθε στοιχείου στην νέα φέτα.
package main

import "fmt"

func main() {

	// Δηλώστε μια nil φέτα ακεραίων.
	var numbers []int

	// Προσθέστε αριθμούς στην φέτα.
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i*10)
	}

	// Παρουσιάστε κάθε τιμή.
	for _, number := range numbers {
		fmt.Println(number)
	}

	// Δηλώστε μια φέτα συμβολοσειρών.
	names := []string{"Bill", "Joan", "Jim", "Cathy", "Beth"}

	// Παρουσιάστε κάθε θέση δείκτη και τιμή φέτας.
	for i, name := range names {
		fmt.Printf("Index: %d  Name: %s\n", i, name)
	}

	// Πάρτε μια φέτα των δεικτών 1 και 2.
	slice := names[1:3]

	// Παρουσιάστε την τιμή της νέας φέτας.
	for i, name := range slice {
		fmt.Printf("Index: %d  Name: %s\n", i, name)
	}
}
