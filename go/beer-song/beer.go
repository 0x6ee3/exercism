package beer

import (
	"errors"
	"fmt"
)

const testVersion = 1

const (
	verseTemplate = "%s of beer on the wall, %s of beer.\nTake one down and pass it around, %s of beer on the wall.\n"
	verseLast     = "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n"
	verseNoBeer   = "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"
)

func toBottle(num int) string {
	if num > 1 {
		return fmt.Sprintf("%d bottles", num)
	}
	if num == 1 {
		return "1 bottle"
	}
	return "no more bottles"
}

// Verse ..
func Verse(num int) (string, error) {
	if num < 0 || num > 99 {
		return "", errors.New("Invalid number of beers")
	}

	if num == 1 {
		return verseLast, nil
	}

	if num == 0 {
		return verseNoBeer, nil
	}

	return fmt.Sprintf(verseTemplate, toBottle(num), toBottle(num), toBottle(num-1)), nil
}

// Verses ..
func Verses(from, to int) (string, error) {
	if from < 0 || from > 99 || to < 0 || to > 99 || from < to {
		return "", errors.New("Invalid number of beers")
	}

	var song string
	for i := from; i >= to; i-- {
		verse, err := Verse(i)
		if err != nil {
			return "", err
		}
		song += verse + "\n"
	}

	return song, nil
}

// Song ..
func Song() string {
	song, _ := Verses(99, 0)
	return song
}
