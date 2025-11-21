package fibonacci

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
func fibonacci(n int) int {
	// Alapesetek (bázis feltételek)
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	
	// Rekurzió: az előző kettő összege
	return fibonacci(n-1) + fibonacci(n-2)
}
