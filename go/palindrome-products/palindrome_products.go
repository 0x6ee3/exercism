package palindrome

const testVersion = 1

// Product ..
type Product struct {
	Product        int
	Factorizations [][2]int
}

// Products ..
func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	return Product{}, Product{}, nil
}
