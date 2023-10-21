//go:build OMIT

package main

import (
	"fmt"
	"math"
	"math/bits"
)

func main() {
	inputs := [][]byte{
		[]byte{0x7F},
		[]byte{0x81, 0x00},
		[]byte{0xC0, 0x00},
		[]byte{0xFF, 0x7F},
		[]byte{0x81, 0x80, 0x00},
		[]byte{0xFF, 0xFF, 0x7F},
		[]byte{0x81, 0x80, 0x80, 0x00},
		[]byte{0xC0, 0x80, 0x80, 0x00},
		[]byte{0xFF, 0xFF, 0xFF, 0x7F},
		[]byte{0x82, 0x00},
		[]byte{0x81, 0x10},
	}

	for _, input := range inputs {
		decoded, err := DecodeVarint(input)
		if err != nil {
			fmt.Println(err)
			return
		}

		encoded := EncodeVarint(decoded)
		fmt.Printf("input 0x%x, decoded: %d, encoded: 0x%x\n", input, decoded, encoded)
	}
}

// Η DecodeVarint παίρνει έναν ακέραιο βασισμένο σε κωδικοποίηση
// μεταβλητού μήκους VLQ and τον αποκωδικοποιεί σε έναν ακέραιο
// με 32 bit.
func DecodeVarint(input []byte) (uint32, error) {
	const lastBitSet = 0x80 // 1000 0000

	var d uint32
	var bitPos int

	for i := len(input) - 1; i >= 0; i-- {
		n := uint8(input[i])

		// Επεξεργαστήτε τα πρώτα 7 bit και αγνοήστε το 8ο.
		for checkBit := 0; checkBit < 7; checkBit++ {

			// Απομακρύνετε περιστρέφοντας το τελευταίο bit και μετακινείστε το στο τέλος.
			// Πριν: 0000 0001
			// Μετά: 1000 0000
			n = bits.RotateLeft8(n, -1)

			// Υπολογείστε με βάση μόνο εκείνα τα bit τιμής 1 που περιστράφηκαν.
			// Μετατρέψτε την bitPos σε αριθμιση με βάση το 10.
			if n >= lastBitSet {
				switch {
				case bitPos == 0:
					d++
				default:
					base10 := math.Pow(2, float64(bitPos))
					d += uint32(base10)
				}
			}

			// Μετακινείστε την θέση bit.
			bitPos++
		}
	}

	return d, nil
}

// Η EncodeVarint παίρνει έναν ακέραιο 32 bit και τον κωδικοποιεί
// σε ένα ακέραιο βασισμένο σε κωδικοποίηση μεταβλητού μήκους VLQ.
func EncodeVarint(n uint32) []byte {
	const maxBytes = 4
	const eightBitSet = 0x80      // 1000 0000
	const lastBitSet = 0x80000000 // 1000 0000 0000 0000

	encoded := make([]byte, maxBytes)

	for bytePos := maxBytes - 1; bytePos >= 0; bytePos-- {
		var d uint8

		// Επεξεργαστείτε τα επόμενα 7 bit.
		for checkBit := 0; checkBit < 7; checkBit++ {

			// Απομακρύνετε το τελευταίο bit περιστρέφοντας το
			// και μετακινήστε το προς το τέλος.
			// Πριν: 0000 0000 0000 0001
			// Μετά: 1000 0000 0000 0000
			n = bits.RotateLeft32(n, -1)

			// Υπολογείστε μόνο με βάση εκείνων των bit τιμής 1
			// που περιστράφηκαν. Μετατρέψτε την θέση bit σε
			// αριθμητική βάσης 10.
			if n >= lastBitSet {
				switch {
				case checkBit == 0:
					d++
				default:
					base10 := math.Pow(2, float64(checkBit))
					d += uint8(base10)
				}
			}
		}

		// Αυτές οι τιμές χρειάζονται το 8ο bit να έχει τιμη ίση με 1.
		if bytePos < 3 {
			d += eightBitSet
		}

		// Αποθηκεύστε την τιμή με αντίστροφη σειρά.
		encoded[bytePos] = d
	}

	// Απομακρύνετε τα αρχικά μηδενικά βρίσκοντας τιμές
	// που έχουν το όγδοο bit με τιμή 1.
	for bytePos, b := range encoded {
		if b == eightBitSet {
			continue
		}
		encoded = encoded[bytePos:]
		break
	}

	return encoded
}
