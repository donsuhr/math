// Package math provides simple integer arithmetic helpers.
package math

import (
	"golang.org/x/exp/constraints"
)

// Number is a constraint satisfied by any integer or floating-point type.
type Number interface {
	constraints.Float | constraints.Integer
}

// Add returns the sum of a and b.
//
// See https://mathsisfun.com/numbers/addition.html for more on addition.
func Add[T Number](a T, b T) T {
	return a + b
}
