//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://github.com/dominikh/go-tools#installation
// go install honnef.co/go/tools/cmd/...@2023.1.3

// Alignment is about placing fields on address alignment boundaries
// for more efficient reads and writes to memory.

// Sample program to show how struct types align on boundaries.
package main

import (
	"fmt"
	"unsafe"
)

// No byte padding.
type nbp struct {
	a bool // 	1 byte				sizeof 1
	b bool // 	1 byte				sizeof 2
	c bool // 	1 byte				sizeof 3 - Aligned on 1 byte
}

// Single byte padding.
type sbp struct {
	a bool //	1 byte				sizeof 1
	//			1 byte padding		sizeof 2
	b int16 // 	2 bytes				sizeof 4 - Aligned on 2 bytes
}

// Three byte padding.
type tbp struct {
	a bool //	1 byte				size 1
	//			3 bytes padding		size 4
	b int32 //	4 bytes				size 8 - Aligned on 4 bytes
}

// Seven byte padding.
type svnbp struct {
	a bool //	1 byte				size 1
	//			7 bytes padding		size 8
	b int64 //	8 bytes				size 16 - Aligned on 8 bytes
}

// No padding.
type np struct {
	a string // 16 bytes			size 16
	b string // 16 bytes			size 32
	c int32  //  4 bytes			size 36
	d int32  //  4 bytes			size 40 - Aligned on 8 bytes
}

// Eight byte padding on 64bit Arch. Word size is 8 bytes.
type ebp64 struct {
	a string //	16 bytes			size 16
	b int32  //	 4 bytes			size 20
	//  		 4 bytes padding	size 24
	c string //	16 bytes			size 40
	d int32  //	 4 bytes			size 44
	//  		 4 bytes padding	size 48 - Aligned on 8 bytes
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
