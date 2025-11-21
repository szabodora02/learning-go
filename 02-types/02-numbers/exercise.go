package calculator

import (
	"math"
	"strconv"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE

// 1. amean: Kiszámolja az átlagot és a legközelebbi egészre kerekíti
func amean(x, y float64) int {
	mean := (x + y) / 2.0
	return int(math.Round(mean))
}

// 2. ameanString: Stringeket vár, átalakítja őket, és meghívja a fenti függvényt
func ameanString(x, y string) (int, error) {
	// Első szám konvertálása
	f1, err := strconv.ParseFloat(x, 64)
	if err != nil {
		return 0, err
	}

	// Második szám konvertálása
	f2, err := strconv.ParseFloat(y, 64)
	if err != nil {
		return 0, err
	}

	// Eredmény kiszámolása
	return amean(f1, f2), nil
}
