package queenattack

import "errors"

const testVersion = 1

// CanQueenAttack ...
func CanQueenAttack(w, b string) (bool, error) {
	if len(w) != 2 || len(b) != 2 ||
		w[0] < 'a' || w[0] > 'h' || w[1] < '1' || w[1] > '8' ||
		b[0] < 'a' || b[0] > 'h' || b[1] < '1' || b[1] > '8' ||
		b == w {
		return false, errors.New("Error ..")
	}

	if w[0] == b[0] || w[1] == b[1] {
		return true, nil
	}

	if abs(int(w[0]-'a')-int(b[0]-'a')) == abs(int(w[1]-'1')-int(b[1]-'1')) {
		return true, nil
	}

	return false, nil
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
