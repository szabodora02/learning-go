package wordcount

import (
	"strings"
	"sync"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE

// CountWords counts words concurrently using goroutines and channels
func CountWords(lines []string) map[string]int {
	// Eredmény map létrehozása
	counts := make(map[string]int)
	
	// Csatorna a szavak továbbítására
	wordChan := make(chan string)
	
	var wg sync.WaitGroup

	// 1. Minden sorhoz indítunk egy goroutine-t (Producer)
	for _, line := range lines {
		wg.Add(1)
		go func(l string) {
			defer wg.Done()
			// Szavakra bontás
			words := strings.Fields(l)
			for _, w := range words {
				wordChan <- w // Beküldjük a szót a közös csatornába
			}
		}(line)
	}

	// 2. Egy külön szálon figyeljük, mikor végez mindenki, és lezárjuk a csatornát
	go func() {
		wg.Wait()
		close(wordChan)
	}()

	// 3. Fő szál: Adatok fogadása és összegzése (Consumer)
	// Ez addig fut, amíg a csatornát le nem zárják
	for w := range wordChan {
		counts[w]++
	}

	return counts
}
