package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	proverbs := make([]string, 0)
	if len(rhyme) == 0 {
		return proverbs
	}

	for i := 1; i < len(rhyme); i++ {
		proverbs = append(proverbs, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i-1], rhyme[i]))
	}
	proverbs = append(proverbs, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	return proverbs
}
