package pipeline

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE

// generator: Slice -> Channel
// Létrehoz egy csatornát, elindít egy goroutine-t a küldéshez, és visszaadja a csatornát.
func generator(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out) // Fontos: lezárjuk, ha végeztünk, hogy a köv. lépés tudja, nincs több adat
	}()
	return out
}

// adder: Channel (int) -> Channel (float32)
// Olvas a bemeneti csatornából, hozzáad 1-et, konvertál, és továbbküld.
func adder(in <-chan int) <-chan float32 {
	out := make(chan float32)
	go func() {
		// A range addig fut, amíg az 'in' csatornát le nem zárják
		for n := range in {
			result := float32(n + 1)
			out <- result
		}
		close(out)
	}()
	return out
}

// collector: Channel (float32) -> Slice
// Összegyűjti az adatokat a csatornából egy tömbbe.
func collector(in <-chan float32) []float32 {
	var result []float32
	// Itt nem kell goroutine, mert ez a "nyelő" (sink), itt várjuk meg a végét.
	for n := range in {
		result = append(result, n)
	}
	return result
}
