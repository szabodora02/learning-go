package strings

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
// 1. Multiline string
func multilineString() string {
	return `some
multiline
string`
}

// 2. String length
func stringLen(s string) int {
	return len(s)
}

// 3. Trim first char (remove the first character)
func trimFirstChar(s string) string {
	// Rúnává (karakterré) alakítjuk, hogy biztosan kezelje az ékezeteket is
	r := []rune(s)
	if len(r) == 0 {
		return ""
	}
	return string(r[1:])
}

// 4. Trim last char (remove the last character)
func trimLastChar(s string) string {
	r := []rune(s)
	if len(r) == 0 {
		return ""
	}
	return string(r[:len(r)-1])
}

// 5. Swap first char with 'A'
func swapFirstChar(s string) string {
	r := []rune(s)
	if len(r) == 0 {
		return ""
	}
	r[0] = 'A'
	return string(r)
}

// 6. Swap last char with 'A'
func swapLastChar(s string) string {
	r := []rune(s)
	if len(r) == 0 {
		return ""
	}
	r[len(r)-1] = 'A'
	return string(r)
}

// 7. Prepend 'A' (Add 'A' to the beginning)
func prependChar(s string) string {
	return "A" + s
}

// 8. Append 'A' (Add 'A' to the end)
func appendChar(s string) string {
	return s + "A"
}
