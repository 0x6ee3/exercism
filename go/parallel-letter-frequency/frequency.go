package letter

// FreqMap ..
type FreqMap map[rune]int

// Frequency ..
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency ..
func ConcurrentFrequency(s []string) FreqMap {
	fmChan := make(chan FreqMap)
	fm := FreqMap{}

	for _, v := range s {
		go func(in string) {
			fmChan <- Frequency(in)
		}(v)
	}

	for _ = range s {
		for r, v := range <-fmChan {
			fm[r] += v
		}
	}

	return fm
}
