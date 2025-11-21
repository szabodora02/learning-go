package sorting

import "sort"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE

// 1. Létrehozunk egy típuskényszert (Constraint), ami felsorolja a megengedett típusokat.
// A feladat kérése: uint8 és int32
type Number interface {
	uint8 | int32
}

// 2. A generikus függvény.
// [T Number] jelzi, hogy a T típusnak meg kell felelnie a fenti szabálynak.
func sortSlice[T Number](input []T) []T {
	// A standard sort.Slice függvényt használjuk, ami helyben rendez.
	// Mivel a T típusok (uint8, int32) összehasonlíthatók a < jellel, ez működni fog.
	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})

	return input
}
