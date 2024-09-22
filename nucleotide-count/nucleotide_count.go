package dna

import (
	"fmt"
)

type Histogram map[rune]uint64

type DNA string

func (d DNA) GetEmpytiHistogram() Histogram {
	return Histogram{
		'A': 0, 'C': 0, 'G': 0, 'T': 0,
	}
}

func (d DNA) Counts() (Histogram, error) {
	h := d.GetEmpytiHistogram()
	for _, nucleotide := range d {
		if _, ok := h[nucleotide]; !ok {
			return nil, fmt.Errorf("invalid nucleotide %v", nucleotide)
		}
		h[nucleotide]++
	}
	return h, nil
}
