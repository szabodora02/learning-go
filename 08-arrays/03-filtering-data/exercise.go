package filteringdata

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// filterData filters a slice based in an index slice.
func filterData(keys []string, indices []int) [10]string {
	var result [10]string

	// 1. szabály: Ha a hosszak nem egyeznek, üres tömbbel térünk vissza
	if len(keys) != len(indices) {
		return result
	}

	// Számláló, hogy hova írjuk a következő elemet az eredménytömbben
	count := 0

	for i := 0; i < len(keys); i++ {
		// 2. szabály: Csak akkor tartjuk meg, ha az index NEM nagyobb mint 5 (<= 5)
		if indices[i] <= 5 {
			// Biztonsági ellenőrzés: ne írjunk túl a 10 elemű tömbön
			if count < 10 {
				result[count] = keys[i]
				count++
			}
		}
	}

	// 3. szabály: A maradék helyek üresek maradnak (a Go-ban alapból üresek a tömb létrehozásakor)
	return result
}
