package summultiples

// MultipleSummer ...
func MultipleSummer(divisor ...int) func(int) int {
	return func(lim int) int {
		var sum int

		for i := 1; i < lim; i++ {
			for _, div := range divisor {
				if i%div == 0 {
					sum += i
					break
				}
			}
		}

		return sum
	}
}
