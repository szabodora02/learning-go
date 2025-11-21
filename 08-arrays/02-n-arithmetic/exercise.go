package narithmetic

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// nArithmetic returns the result of an arithmetic operation over "n" elements.
func nArithmetic(elems [10]int) int {
	// Az első elemmel kezdünk (ez a kiindulási alap)
	result := elems[0]

	// A 2. elemtől (ami a 1-es index) megyünk végig a tömb végéig
	for i := 1; i < len(elems); i++ {
		result -= elems[i]
	}

	return result
}
