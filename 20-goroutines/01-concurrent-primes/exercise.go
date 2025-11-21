package concurrentprimes

import (
	"math"
	"sort"
	"sync"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// Segédfüggvény: Eldönti egy számról, hogy prím-e
func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	limit := int(math.Sqrt(float64(num)))
	for i := 2; i <= limit; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// GeneratePrimes returns primes <= n using concurrency
func GeneratePrimes(n int) []int {
	// FONTOS: Ha n < 2, azonnal visszaadunk egy üres (de nem nil!) szeletet
	if n < 2 {
		return []int{}
	}

	var wg sync.WaitGroup
	ch := make(chan int)

	// 1. Indítjuk a goroutine-okat
	for i := 2; i <= n; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			if isPrime(val) {
				ch <- val
			}
		}(i)
	}

	// 2. Lezárás külön szálon
	go func() {
		wg.Wait()
		close(ch)
	}()

	// 3. Begyűjtés
	// FONTOS: make-kel hozzuk létre, hogy biztosan ne legyen nil, még ha üres is marad
	primes := make([]int, 0)
	for p := range ch {
		primes = append(primes, p)
	}

	// 4. Sorbarendezés
	sort.Ints(primes)

	return primes
}
