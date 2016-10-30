package sieve

// Sieve algorithm (without optimization)
func Sieve(n int) []int {

	arr := make(map[int]bool, n)
	for i := 2; i < n; i++ {
		if !arr[i] {
			for j := i * i; j < n; j += i {
				arr[j] = true
			}
		}
	}

	var prim []int
	for i := 2; i < n; i++ {
		if !arr[i] {
			prim = append(prim, i)
		}
	}

	return prim
}
