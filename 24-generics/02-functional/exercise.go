package functional

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE

// flatten takes a slice of slices and merges them into a single slice.
// [T any] jelenti, hogy a függvény bármilyen típussal működik.
// Bemenet: [][]T (pl. [[1,2], [3,4]])
// Kimenet: []T (pl. [1,2,3,4])
func flatten[T any](lists [][]T) []T {
	var result []T

	// Végigmegyünk a külső listán (ami belső listákat tartalmaz)
	for _, innerList := range lists {
		// A három pont (...) "kicsomagolja" a belső lista elemeit,
		// így az append hozzá tudja őket fűzni az eredményhez.
		result = append(result, innerList...)
	}

	return result
}
