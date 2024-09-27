package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	if shiftKey < 0 {
		return ""
	}

	shiftKey = shiftKey % 26
	cipher := make([]rune, len(plain))

	for i, char := range plain {
		switch {
		case char >= 'A' && char <= 'Z':
			cipher[i] = 'A' + ((char - 'A' + rune(shiftKey)) % 26) // Maiúsculas
		case char >= 'a' && char <= 'z':
			cipher[i] = 'a' + ((char - 'a' + rune(shiftKey)) % 26) // Minúsculas
		default:
			cipher[i] = char
		}
	}

	return string(cipher)
}
