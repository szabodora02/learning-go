package logicalops
// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
// 1. Inverse returns the logical negation (NOT)
func inverse(b bool) bool {
	return !b
}

// 2. And returns the logical AND
func and(x, y bool) bool {
	return x && y
}

// 3. DeMorgan returns NOT (A OR B) using De Morgan's laws
// Formula: (NOT a) AND (NOT b)
// Constraint: Use only the 'inverse' and 'and' functions defined above!
func deMorgan(a, b bool) bool {
	// Így kell összerakni a fenti függvényekből:
	return and(inverse(a), inverse(b))
}
