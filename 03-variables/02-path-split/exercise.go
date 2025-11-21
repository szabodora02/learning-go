package pathsplit

import "path"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
// splitPath returns the file component of a file path.
func splitPath(fullPath string) string {
	// A path.Split visszaadja a könyvtárat és a fájlt.
	// A könyvtár rész nem kell, ezért "_" jellel eldobjuk.
	_, file := path.Split(fullPath)
	return file
}
