package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(calc FodderCalculator, cows int) (float64, error) {
	totalFodder, err := calc.FodderAmount(cows)
	if err != nil {
		return 0.0, err
	}
	factor, err := calc.FatteningFactor()
	if err != nil {
		return 0.0, err
	}

	return float64(totalFodder*factor) / float64(cows), nil
}

func ValidateInputAndDivideFood(calc FodderCalculator, cows int) (float64, error) {
	if cows > 0 {
		return DivideFood(calc, cows)
	}

	return 0.0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
	cows    int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{
			cows:    cows,
			message: "there are no negative cows",
		}
	}

	if cows == 0 {
		return &InvalidCowsError{
			cows:    cows,
			message: "no cows don't need food",
		}
	}

	return nil
}
