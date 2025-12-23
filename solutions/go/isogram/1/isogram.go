package isogram

import (
    "strings"
    "fmt"
)

func IsIsogram(word string) bool {
	if word == "" {
        return true
    }
    
	count := 0
    for _, v := range word {
        l := fmt.Sprintf("%c", v)
		if l == " " || l == "-" {
            continue
        }
        c := strings.Count(strings.ToLower(word), strings.ToLower(l))
        if c > count {
            count = c
        }
    }
    return count <= 1
}
