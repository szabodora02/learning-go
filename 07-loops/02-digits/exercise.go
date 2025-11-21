package digits

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
func sumDigits(n int) int {
	sum := 0
	// Addig megyünk, amíg van számjegy (n > 0)
	for n > 0 {
		// 1. Megszerezzük az utolsó számjegyet (maradékos osztás 10-zel)
		digit := n % 10
		
		// 2. Hozzáadjuk az összeghez
		sum += digit
		
		// 3. Levágjuk az utolsó számjegyet (egészosztás 10-zel)
		n = n / 10
	}
	return sum
}
