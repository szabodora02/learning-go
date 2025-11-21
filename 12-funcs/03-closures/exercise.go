package closures

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate
func proxy(funcs ...func() int) func() int {
	// Ez a változó a "closure state". 
	// A visszaadott függvény "emlékezni" fog rá minden hívásnál.
	current := 0

	// Visszaadunk egy névtelen függvényt (ez maga a closure)
	return func() int {
		// Biztonsági ellenőrzés: ha üres listát kaptunk
		if len(funcs) == 0 {
			return 0
		}

		// 1. Kiválasztjuk a soron következő függvényt
		selectedFunc := funcs[current]

		// 2. Kiszámoljuk az új indexet a következő híváshoz.
		// A % (modulo) operátor biztosítja a körbeforgást.
		// Pl. ha 3 függvény van: 0->1, 1->2, 2->0
		current = (current + 1) % len(funcs)

		// 3. Lefuttatjuk a kiválasztott függvényt és visszaadjuk az eredményét
		return selectedFunc()
	}
}
