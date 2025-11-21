package sleepSort

import "time"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE

// sleepSort returns the input uint-slice sorted in the forward order using concurrency.
func sleepSort(input []uint) []uint {
	// Csatorna, amin a "felébredt" számok érkeznek
	ch := make(chan uint)

	// 1. Minden számra indítunk egy goroutine-t
	for _, n := range input {
		// FONTOS: 'n'-et paraméterként adjuk át a névtelen függvénynek!
		// Ha nem így tennénk (closure capture), minden szál lehet, hogy ugyanazt az 'n'-t látná.
		go func(val uint) {
			// A feladat szerint: x * 10 millisecond várakozás
			duration := time.Duration(val) * 10 * time.Millisecond
			time.Sleep(duration)
			
			// Ébredés után beküldjük a számot
			ch <- val
		}(n)
	}

	// 2. Összegyűjtjük az eredményeket
	// Mivel tudjuk, hány elem van (len(input)), pontosan annyiszor olvasunk a csatornából.
	var output []uint
	for i := 0; i < len(input); i++ {
		// Ez itt blokkol, amíg a következő leggyorsabb elem meg nem érkezik
		num := <-ch
		output = append(output, num)
	}

	return output
}
