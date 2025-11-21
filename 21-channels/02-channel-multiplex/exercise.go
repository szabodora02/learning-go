package channelmultiplexer

import (
	"context"
	"sync"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// channelMultiplex combines multiple input channels into one output channel.
func channelMultiplex(ctx context.Context, inputs []chan any) chan any {
	out := make(chan any)
	var wg sync.WaitGroup

	// Segédfüggvény, ami egy darab bemeneti csatornát kezel
	multiplex := func(c chan any) {
		defer wg.Done()
		for {
			select {
			// 1. Ha a Contextet törlik, azonnal leállunk a figyeléssel
			case <-ctx.Done():
				return
			// 2. Olvasunk a bemeneti csatornából
			case val, ok := <-c:
				if !ok {
					return // A bemeneti csatornát lezárták
				}
				// 3. Megpróbáljuk elküldeni a kimenetre, DE figyeljük a contextet is
				select {
				case out <- val:
					// Sikeres küldés
				case <-ctx.Done():
					return // Küldés közben törölték a contextet
				}
			}
		}
	}

	// Minden bemeneti csatornához indítunk egy goroutine-t
	for _, ch := range inputs {
		wg.Add(1)
		go multiplex(ch)
	}

	// Egy külön szálon várjuk, hogy mindenki végezzen, majd lezárjuk a kimenetet
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
