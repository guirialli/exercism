package strand

import "errors"

type RNA struct {
	Sequence string
}

func NewRNAByDNA(dnaSequence string) (*RNA, error) {
	rna := RNA{}
	sequence, err := rna.convertDnaToRNA(dnaSequence)
	if err != nil {
		return nil, err
	}
	return &RNA{Sequence: sequence}, nil
}

func (r *RNA) convertDnaToRNA(dna string) (string, error) {
	sequence := ""
	for _, char := range dna {
		p, err := r.partDnaToRna(char)
		if err != nil {
			return "", err
		}
		sequence += string(p)
	}
	return sequence, nil
}

func (r *RNA) partDnaToRna(dna rune) (rune, error) {
	switch dna {
	case 'G':
		return 'C', nil
	case 'C':
		return 'G', nil
	case 'T':
		return 'A', nil
	case 'A':
		return 'U', nil
	default:
		return 0, errors.New("invalid DNA")
	}
}

func ToRNA(dna string) string {
	rna, err := NewRNAByDNA(dna)
	if err != nil {
		panic(err)
	}
	return rna.Sequence
}
