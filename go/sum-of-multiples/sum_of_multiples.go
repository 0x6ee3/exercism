package summultiples

// MultipleSummer ...
func MultipleSummer(divisor ...int) func(int) int {
	return func(in int) int {

		values := make(map[int]bool)
		for _, div := range divisor {
			for i := 1; i*div < in; i++ {
				values[div*i] = true
			}
		}

		var sum int
		for v := range values {
			sum += v
		}

		return sum
	}
}
