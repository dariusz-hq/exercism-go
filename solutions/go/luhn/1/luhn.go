package luhn

import (
    "unicode/utf8"
    "strings"
    "strconv"
)

func Valid(id string) bool {
	length := utf8.RuneCountInString(id)
	if length == 1 {
		return false
	}

	slice := strings.Split(id, "")
	reversed := []string{}
	for i := length - 1; i >= 0; i-- {
		if slice[i] != " " {
			reversed = append(reversed, slice[i])
		}
	}
    if len(reversed) == 1 {
        return false
    }

	res := []int{}
	for i := 1; i <= len(reversed); i++ {
		strV := reversed[i-1]
		intV, err := strconv.Atoi(strV)
        if err != nil {
            return false
        }
		if i%2 == 0 {
			prod := intV * 2
			if prod > 9 {
				res = append(res, prod-9)
			} else {
				res = append(res, prod)
			}
		} else {
			res = append(res, intV)
		}
	}

	sum := 0
	for _, v := range res {
		sum += v
	}

	return sum%10 == 0
}
