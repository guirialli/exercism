package series

func All(n int, s string) []string {
	serie := make([]string, 0)
	for i := 0; i+n <= len(s); i++ {
		serie = append(serie, s[i:i+n])
	}
	return serie
}

func UnsafeFirst(n int, s string) string {
	return s[0:n]
}
