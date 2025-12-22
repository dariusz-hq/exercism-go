package parsinglogfiles

import (
    "regexp"
    "fmt"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`(^\[TRC\])|(^\[DBG\])|(^\[INF\])|(^\[WRN\])|(^\[ERR\])|(^\[FTL\])`)
    if re.MatchString(text) {
        return true
    }
    return false
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~=*-]*>`)
    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)
    count := 0
    for _, v := range lines {
        matches := re.FindStringSubmatch(v)
        count += len(matches)
    }
    return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)
    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+([a-zA-Z1-9]+)`)
    result := []string{}
    for _, v := range lines {
        matches := re.FindStringSubmatch(v)
        prefix := ""
        if len(matches) > 0 {
            prefix = fmt.Sprintf("[USR] %s ", matches[1])
        }
        line := fmt.Sprintf("%s%s", prefix, v)
        result = append(result, line)
    }
    return result
}
