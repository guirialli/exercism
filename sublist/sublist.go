package sublist

import "sort"

// Relation type is defined in relations.go file.

type sublist struct {
	l1 []int
	l2 []int
}

func NewSublist(l1, l2 []int) *sublist {
	sort.Ints(l2)
	sort.Ints(l1)
	if len(l2) > len(l1) {
		return &sublist{
			l1: l2,
			l2: l1,
		}
	}
	return &sublist{
		l1: l1,
		l2: l2,
	}
}

func (s *sublist) IsEqual() bool {
	if len(s.l1) != len(s.l2) {
		return false
	}
	for i := 0; i < len(s.l1); i++ {
		if s.l1[i] != s.l2[i] {
			return false
		}
	}
	return true
}
func (s *sublist) IsSublist() bool {
	if len(s.l1) == 0 && len(s.l2) != 0 || len(s.l2) == 0 && len(s.l1) != 0 {
		return true
	}
	for _, e2 := range s.l2 {
		for _, e1 := range s.l1 {
			if e2 != e1 {
				return false
			}
		}
	}
	return true
}

func (s *sublist) Relation() Relation {
	if s.IsEqual() {
		return RelationEqual
	} else if s.IsSublist() {
		return RelationSublist
	}
	return RelationUnequal
}

func Sublist(l1, l2 []int) Relation {
	return NewSublist(l1, l2).Relation()
}
