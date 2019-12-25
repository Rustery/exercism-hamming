// Package hamming is Exercism.io exercise
package hamming

import (
	"errors"
	"unicode/utf8"
)

// Distance — Calculating Hamming Distance for two DNA strands
func Distance(a, b string) (int, error) {

	if utf8.RuneCountInString(a) != utf8.RuneCountInString(b) {
		return 0, errors.New("strands should be equal size")
	}

	var hd int

	for i := 0; i < utf8.RuneCountInString(a); i++ {
		if a[i] != b[i] {
			hd++
		}
	}

	return hd, nil
}
