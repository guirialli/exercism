package protein

import (
	"errors"
)

var ErrStop = errors.New("stop")
var ErrInvalidBase = errors.New("invalid base")

type Protein map[string]string

func New() Protein {
	return Protein{
		"UUU": "Phenylalanine", "UUC": "Phenylalanine", "UUA": "Leucine", "UUG": "Leucine", "UCU": "Serine",
		"UCC": "Serine", "UCA": "Serine", "UCG": "Serine", "UAU": "Tyrosine", "UAC": "Tyrosine", "UGU": "Cysteine",
		"UGC": "Cysteine", "UGG": "Tryptophan", "UAA": "STOP", "AUG": "Methionine", "UAG": "STOP", "UGA": "STOP",
	}
}

func FromRNA(rna string) ([]string, error) {
	result := make([]string, 0)
	if len(rna)%3 != 0 {
		return result, ErrInvalidBase
	}

	protens := New()
	for i := 0; i < len(rna)/3; i++ {
		codon := rna[i*3 : (i+1)*3]
		if codon == "AUG" && i != 0 || protens[codon] == "STOP" {
			break
		}
		result = append(result, protens[codon])

	}
	return result, nil
}

func FromCodon(codon string) (string, error) {
	c, ok := New()[codon]
	if !ok {
		return "", ErrInvalidBase
	} else if c == "STOP" {
		return "", ErrStop
	}
	return c, nil
}
