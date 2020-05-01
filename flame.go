package flame

import (
	"fmt"
	"strings"
)

var flameMap = map[string]string{
	"F": "Friends",
	"L": "Lovers",
	"A": "Admirer",
	"M": "Married",
	"E": "Enemy",
}

func removeCharacters(input string, characters string) string {
	println("HELLO")
	input = strings.ToLower(input)
	characters = strings.ToLower(characters)

	filter := func(r rune) rune {
		if strings.IndexRune(characters, r) < 0 {
			return r
		}
		return -1
	}

	return strings.Map(filter, input)
}

// Pair returns a concatenated result of name1 and name2.
func Pair(name1, name2 string) string {
	var x, y string

	for _, rune1 := range name1 {
		for _, rune2 := range name2 {
			if rune1 == rune2 {
				x = removeCharacters(name1, string(rune1))
				y = removeCharacters(name2, string(rune2))
			}
		}
	}

	return fmt.Sprintf("%s%s", x, y)
}

// GetResult returns the `FLAME` value of the pair.
func GetResult(pair string) string {
	var FLAME = "FLAME"
	var flameLength = 5
	var index int

	for len(FLAME) != 1 {
		index = len(pair) % flameLength

		if index == 0 {
			FLAME = removeCharacters(FLAME, string(FLAME[len(FLAME)-1]))
		} else {
			FLAME = removeCharacters(FLAME, string(FLAME[index-1]))
			FLAME = FLAME[index-1:len(FLAME)] + FLAME[0:index-1]
		}

		flameLength--
	}
	return flameMap[strings.ToUpper(FLAME)]
}
