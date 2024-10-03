package sublist

import "reflect"

// Relation type is defined in relations.go file.

func isSubilist(l1, l2 []int) bool {
	if len(l1) < len(l2) {
		return false
	}
	for i := 0; i <= len(l1)-len(l2); i++ {
		if reflect.DeepEqual(l1[i:i+len(l2)], l2) {
			return true
		}
	}
	return false
}
func Sublist(l1, l2 []int) Relation {
	switch {
	case reflect.DeepEqual(l1, l2):
		return RelationEqual
	case isSubilist(l2, l1):
		return RelationSublist
	case isSubilist(l1, l2):
		return RelationSuperlist
	default:
		return RelationUnequal
	}
}
