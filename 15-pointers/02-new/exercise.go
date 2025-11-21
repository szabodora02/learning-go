package pointernew

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate
func newValue() *bool {
	// A new(bool) lefoglal egy bool-t, beállítja false-ra (zero value),
	// és visszaadja a pointerét (*bool). Pont ez kell nekünk.
	return new(bool)
}
