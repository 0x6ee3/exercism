package diffsquares

// SquareOfSums ...
func SquareOfSums(n int) int {
	var sum int

	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum * sum
}

// SumOfSquares ...
func SumOfSquares(n int) int {
	var sum int

	for i := 1; i <= n; i++ {
		sum += i * i
	}

	return sum
}

// Difference ...
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
