package logs

import (
    "fmt"
    "unicode/utf8"
)

const (
    Recommendation string = "U+2757"
    Search string = "U+1F50D"
    Weather string = "U+2600"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
    app := ""
    foundCode := ""
	for _, v := range log {
        code := fmt.Sprintf("%U", v)
        if code == Recommendation || code == Search || code == Weather {
            foundCode = code
            break
        }
    }
    switch foundCode {
    case Recommendation:
        app = "recommendation"
    case Search:
        app = "search"
    case Weather:
        app = "weather"
    default:
        app = "default"
    }
    return app
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	newLog := ""

    for _, v := range log {
        r := fmt.Sprintf("%c", v)
        nr := r
        if r == fmt.Sprintf("%c", oldRune) {
            nr = fmt.Sprintf("%c", newRune)
        }
        newLog = fmt.Sprintf("%s%s", newLog, nr)
    }

    return newLog
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	if utf8.RuneCountInString(log) <= limit {
        return true
    }
    return false
}
