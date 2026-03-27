package thefarm

import (
    "errors"
    "fmt"
)

func DivideFood(fc FodderCalculator, n int) (float64, error) {
    fodderAmt, err := fc.FodderAmount(n)
    if err != nil {
        return 0.0, err
    }
    fatFactor, err := fc.FatteningFactor()
    if err != nil {
        return 0.0, err
    }
    return fodderAmt / float64(n) * fatFactor, nil   
}

func ValidateInputAndDivideFood(fc FodderCalculator, n int) (float64, error) {
    if n > 0 {
        return DivideFood(fc, n)
    }
    return 0.0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
    numberOfCows int
    customMessage string
}

func (err InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", err.numberOfCows, err.customMessage)
}

func ValidateNumberOfCows(n int) error {
	if n < 0 {
        return &InvalidCowsError{
            n,
            "there are no negative cows",
        }
    } else if n == 0 {
        return &InvalidCowsError{
            n,
            "no cows don't need food",
        }
    } else {
        return nil
    }
}

