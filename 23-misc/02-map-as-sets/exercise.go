package search

import (
	"io"
	"regexp"
	"strings"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// contain reads a text and a word and returns true if the word appears in the text MORE THAN ONCE.
func contain(reader io.Reader, word string) bool {
	// 1. Beolvassuk a teljes szöveget
	content, _ := io.ReadAll(reader)
	text := string(content)

	// 2. Kisbetűssé alakítjuk a szöveget és a keresett szót is
	text = strings.ToLower(text)
	targetWord := strings.ToLower(word)

	// 3. Eltávolítjuk az írásjeleket regexp segítségével
	// A [^a-z0-9 ]+ jelentése: minden, ami NEM betű, szám vagy szóköz.
	reg := regexp.MustCompile("[^a-z0-9 ]+")
	text = reg.ReplaceAllString(text, "")

	// 4. Szavakra bontjuk
	words := strings.Fields(text)

	// 5. Megszámoljuk az előfordulásokat egy map segítségével
	counts := make(map[string]int)
	for _, w := range words {
		counts[w]++
	}

	// 6. Ellenőrizzük a feltételt: "appears more than once" (tehát > 1)
	return counts[targetWord] > 1
}
