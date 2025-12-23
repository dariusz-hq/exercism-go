package collatzconjecture

import "errors"

func even(n int) bool {
    return n % 2 == 0
}

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
        return 0, errors.New("N must be a positive number")
    }
	steps := 0
    newNumber := n
    for newNumber != 1 {
        if even(newNumber) == true {
            newNumber = newNumber / 2
        } else {
            newNumber = (newNumber * 3) + 1
        }
        steps++
    }
    return steps, nil
}
