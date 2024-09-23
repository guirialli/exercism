package reverse

func Reverse(input string) string {
	inputRunes := []rune(input)
	reverseInput := ""
	for i := len(inputRunes) - 1; i >= 0; i-- {
		reverseInput += string(inputRunes[i])
	}
	return reverseInput
}
