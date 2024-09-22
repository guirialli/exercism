package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	pattern := `^\[(TRC|DBG|INF|WRN|ERR|FTL)\]\s.*$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	pattern := `<[~*=-]*>`
	re := regexp.MustCompile(pattern)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	pattern := `"[^"]*?\b(?i)password\b[^"]*?`
	re := regexp.MustCompile(pattern)
	count := 0
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	pattern := `end-of-line\d+`
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	pattern := `User\s+(\w+)`
	re := regexp.MustCompile(pattern)

	taggedLines := []string{}
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if match != nil {
			userName := match[1]
			taggedLines = append(taggedLines, fmt.Sprintf("[USR] %s %s", userName, line))
		} else {
			taggedLines = append(taggedLines, line)
		}
	}
	return taggedLines
}
