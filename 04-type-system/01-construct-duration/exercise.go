package constructduration

import "time"

//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// constructTime constructs a `Time` instant based on its two arguments (arg1, arg2)
func constructDuration(arg1 int, arg2 int) time.Duration {
	// Go-ban a szorzáshoz a típusoknak egyezniük kell.
	// Az 'arg1'-et (órák) és 'arg2'-t (percek) át kell alakítani time.Duration típusra.
	return time.Duration(arg1)*time.Hour + time.Duration(arg2)*time.Minute
}
