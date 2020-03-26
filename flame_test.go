package flame_test

import (
	"testing"

	"github.com/rbo13/flame"
)

// works by removing the common
// character occurence of the 2 names
// then counting the renaming characters.
// the number of characters left, determines
// how do you remove each character from the
// F L A M E character.
//
// E.g: For test case 1, we remove
// both characters `e`. The remaining
// characters are `abcd` the length
// of the remaining characters: 4.
// Thus, we remove the letters from the
// F L A M E string that are position at the
// 4th. Counting is done in a circular manner.

// First Iteration: F L A M E
// Output: F L A E -> M is removed
// Second Iteration: F L A E -> start counting at `E`
// Output: F L E -> A is removed
// Third Iteration: F L E -> start counting at `E`
// Output: F L -> E is removed
// Fourth Iteration: F L -> start counting at `F`
// Output: F

// Result of F L A M E is: F stands for Friends :)
var nameTests = []struct {
	title string
	names []string
	out   string
}{
	{
		title: "Should Be Friends",
		names: []string{"abcde", "e"},
		out:   "Friends",
	},
	{
		title: "Should Be Lovers",
		names: []string{"abcdef", "f"},
		out:   "Lovers",
	},
	{
		title: "Should Be Admirer",
		names: []string{"ab", "ac"},
		out:   "Admirer",
	},
	{
		title: "Should Be Married",
		names: []string{"abcd", "d"},
		out:   "Married",
	},
	{
		title: "Should Be Enemy",
		names: []string{"ab", "a"},
		out:   "Enemy",
	},
}

func TestFlames(t *testing.T) {
	for _, tt := range nameTests {
		t.Run(tt.title, func(t *testing.T) {
			res := flame.Pair(tt.names[0], tt.names[1])
			val := flame.GetResult(res)

			if val != tt.out {
				t.Errorf("got %q, want %q", val, tt.out)
			}
		})
	}
}
