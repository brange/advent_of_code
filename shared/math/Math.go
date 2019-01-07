package math

import "math"

// AbsInt returns the absolute value for the int a
func AbsInt(a int) int {
	return int(math.Abs(float64(a)))
}

// MaxInt returns the max value of a and b
func MaxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

// MinInt returns the min value of a and b
func MinInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
