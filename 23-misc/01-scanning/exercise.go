package scanning

import (
	"io"
	"strings"
	"unicode"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// counter reads a text and returns the number of "lowercase words".
func counter(reader io.Reader) int {
	// 1. Beolvassuk a teljes tartalmat
	// Az io.ReadAll byte slice-ot ad vissza
	content, _ := io.ReadAll(reader)

	// 2. Stringgé konvertáljuk és szavakra bontjuk
	// A strings.Fields a whitespace (szóköz, tab, újsor) mentén darabol
	words := strings.Fields(string(content))

	count := 0
	for _, w := range words {
		// Megvizsgáljuk, hogy a szó "kisbetűs"-e
		isLowerCase := true
		
		for _, r := range w {
			// Ha találunk benne BÁRMILYEN nagybetűt, akkor nem számít annak
			if unicode.IsUpper(r) {
				isLowerCase = false
				break
			}
		}

		// Ha a ciklus végigfutott és nem talált nagybetűt, akkor növeljük a számlálót
		if isLowerCase && len(w) > 0 {
			count++
		}
	}

	return count
}
