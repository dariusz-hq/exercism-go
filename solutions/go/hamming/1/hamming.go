package hamming

import (
    "unicode/utf8"
    "errors"
    "strings"
)

func Distance(a, b string) (int, error) {
	// check if strings have the same length using runes
    if utf8.RuneCountInString(a) != utf8.RuneCountInString(b) {
        return 0, errors.New("Sequences are different")
    }
    distance := 0
	max := utf8.RuneCountInString(a)
    aSlice := strings.Split(a, "")
    bSlice := strings.Split(b, "")

    for i := 0; i < max; i++ {
        if aSlice[i] != bSlice[i] {
            distance++
        }
    }

    return distance, nil
}
