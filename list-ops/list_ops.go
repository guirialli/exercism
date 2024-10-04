package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	accumulator := initial
	for _, v := range s {
		accumulator = fn(accumulator, v)
	}
	return accumulator
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	accumulator := initial
	for i := len(s) - 1; i >= 0; i-- {
		accumulator = fn(s[i], accumulator)
	}
	return accumulator
}

func (s IntList) Filter(fn func(int) bool) IntList {
	filter := make([]int, 0)
	for _, v := range s {
		if fn(v) {
			filter = append(filter, v)
		}
	}
	return filter
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	mapper := make([]int, s.Length())
	for i, v := range s {
		mapper[i] = fn(v)
	}
	return mapper
}

func (s IntList) Reverse() IntList {
	reverse := make([]int, 0)
	for i := len(s) - 1; i >= 0; i-- {
		reverse = append(reverse, s[i])
	}
	return reverse
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	concat := s
	for _, v := range lists {
		concat = append(concat, v...)
	}
	return concat
}
