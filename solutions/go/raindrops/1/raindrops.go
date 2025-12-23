package raindrops

import "fmt"

func DivisibleBy3(n int) bool {
    return n % 3 == 0
}

func DivisibleBy5(n int) bool {
    return n % 5 == 0
}

func DivisibleBy7(n int) bool {
    return n % 7 == 0
}

func Convert(number int) string {
	res := ""
    if DivisibleBy3(number) == true {
        res = "Pling"
    }
    if DivisibleBy5(number) == true {
        res = fmt.Sprintf("%s%s", res, "Plang")
    }
    if DivisibleBy7(number) == true {
        res = fmt.Sprintf("%s%s", res, "Plong")
    }
    if res == "" {
        return fmt.Sprintf("%d", number)
    }
    return res
}
