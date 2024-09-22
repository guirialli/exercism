package strain

type Filter[T any] func(T) bool

func Keep[T any](values []T, filter Filter[T]) []T {
	arr := make([]T, 0)
	for _, v := range values {
		if filter(v) {
			arr = append(arr, v)
		}
	}
	return arr
}

func Discard[T any](values []T, filter Filter[T]) []T {
	arr := make([]T, 0)
	for _, v := range values {
		if !filter(v) {
			arr = append(arr, v)
		}
	}
	return arr
}
