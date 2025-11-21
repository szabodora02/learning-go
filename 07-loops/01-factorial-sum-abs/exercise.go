package factorial

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
func calcSum(n int) int {
	sum := 0
	// Hagyományos for ciklus: i=1-től indul, amíg i <= n, minden körben növeljük 1-gyel
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}
