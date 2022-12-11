package util

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Float | constraints.Integer
}

func Abs[T Number](n T) T {
	if n < 0 {
		return -n
	}

	return n
}
