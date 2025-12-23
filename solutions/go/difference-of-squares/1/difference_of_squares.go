package diffsquares

func square(n int) int {
    return n * n
}

func SquareOfSum(n int) int {
    sum := 0
	for n >= 1 {
        sum += n
        n--
    }
    return square(sum)
}

func SumOfSquares(n int) int {
	sum := 0
    for n >= 1 {
        sum += square(n)
        n--
    }
    return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
