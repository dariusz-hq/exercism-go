package dna

import (
    "errors"
)

var NucleotideMap = map[rune]bool{
    'A': true,
    'C': true,
    'G': true,
    'T': true,
}

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
// Start by uncommenting the following line:
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
// Start by uncommenting the following line:
type DNA []rune

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
//
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	var h Histogram
	if !d.Valid(){
        return nil, errors.New("Invalid histogram")
    }
    h = NewHistogram()
    for _, v := range d {
        h[v]++
    }
	return h, nil
}

func (d DNA) Valid() bool {
    for i := range d {
        _, found := NucleotideMap[d[i]]
        if !found {
            return false
        }
    }
    return true
}

func NewHistogram() Histogram {
    return map[rune]int{'A': 0, 'C': 0, 'G': 0, 'T': 0}
}
