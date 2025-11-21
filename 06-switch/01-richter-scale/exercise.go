package richterscale

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// describeEarthquake returns the "description" of a given magnitude value on the Richter scale.
func describeEarthquake(magnitude float32) string {
	// Feltétel nélküli switch: az első "igaz" eset fut le.
	// Mivel sorban haladunk, elég mindig csak a felső korlátot nézni
	// (pl. ha nem < 2.0, akkor biztosan >= 2.0, így a következőnél elég a < 3.0-t nézni).
	switch {
	case magnitude < 2.0:
		return "micro"
	case magnitude < 3.0:
		return "very minor"
	case magnitude < 4.0:
		return "minor"
	case magnitude < 5.0:
		return "light"
	case magnitude < 6.0:
		return "moderate"
	case magnitude < 7.0:
		return "strong"
	case magnitude < 8.0:
		return "major"
	case magnitude < 10.0:
		return "great"
	default:
		// Minden ami 10.0 vagy nagyobb
		return "massive"
	}
}
