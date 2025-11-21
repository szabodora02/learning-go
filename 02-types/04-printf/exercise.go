package printer

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
import "fmt"

// 1. Boolean formázása: "variable of type boolean and value <true/false>"
func printBool(b bool) string {
	return fmt.Sprintf("variable of type boolean and value %t", b)
}

// 2. Egész szám formázása: "variable of type integer and value <szám>"
func printInt(i int) string {
	return fmt.Sprintf("variable of type integer and value %d", i)
}

// 3. Hexadecimális formázás: "variable of type integer in hexadecimal form and value <hex>"
func printHex(i int) string {
	// %x a kisbetűs hexadecimális formátum
	return fmt.Sprintf("variable of type integer in hexadecimal form and value %x", i)
}

// 4. Lebegőpontos formázás: "variable of type float and value <szám 2 tizedessel>"
func printFloat(f float64) string {
	// %.2f jelenti a 2 tizedesjegy pontosságot
	return fmt.Sprintf("variable of type float and value %.2f", f)
}

// 5. String formázása idézőjelekkel: variable of type string and value "<szöveg>"
func printString(s string) string {
	// \" jelenti az idézőjelet a stringen belül
	return fmt.Sprintf("variable of type string and value \"%s\"", s)
}

// 6. Két string összefűzése
func concatStrings(a, b string) string {
	return a + b
}

// 7. Összefűzés és formázott kiíratás (újrahasznosítjuk a fenti függvényeket)
func printConcatStrings(a, b string) string {
	concatenated := concatStrings(a, b)
	return printString(concatenated)
}
