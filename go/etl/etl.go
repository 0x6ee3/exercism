package etl

// Transform ...
func Transform(in map[int][]string) (out map[string]int) {
	out = make(map[string]int)

	var value int
	var key string
	var chars []string

	for value, chars = range in {
		for _, key = range chars {
			// improves performance by 25% compared to strings.ToLower()
			// only works for single byte characters
			out[string(key[0]+32)] = value // 'A' + 32 = 'a'
		}
	}

	return
}
