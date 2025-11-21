package pointerbasic

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate
func retrieveValue(pointer *bool) bool {
	// A * jel a változó neve előtt olvassa ki az értéket a memóriacímről
	return *pointer
}
