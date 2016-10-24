package raindrops

import "strconv"

const testVersion = 2

// Convert a number to the lovely sound of raindrops.
func Convert(rain int) (sound string) {
	if rain%3 == 0 {
		sound = "Pling"
	}
	if rain%5 == 0 {
		sound += "Plang"
	}
	if rain%7 == 0 {
		sound += "Plong"
	}
	if len(sound) == 0 {
		// ~40% faster than fmt.Sprintf("%v", rain)
		return strconv.Itoa(rain)
	}

	return
}
