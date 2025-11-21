package subtask

import (
	"context"
	"time"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE

// StartTask manages the lifecycle of the SubTask
func StartTask(ctx context.Context) (string, error) {
	// 1. Létrehozunk egy új contextet 1 másodperces időkorláttal
	// A defer cancel() kötelező, hogy felszabadítsuk az erőforrásokat
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	// Csatorna az eredmény fogadására a goroutine-ból
	// Egy structot használunk, hogy a stringet és az errort is át tudjuk küldeni
	type result struct {
		val string
		err error
	}
	ch := make(chan result)

	// 2. Elindítjuk a SubTask-ot egy külön goroutine-ban
	go func() {
		// Fontos: lezárjuk a csatornát a végén, ahogy a feladat kéri
		defer close(ch)
		
		res, err := SubTask(ctx)
		ch <- result{val: res, err: err}
	}()

	// Várakozunk: vagy a context jár le (timeout/cancel), vagy megjön az eredmény
	select {
	case <-ctx.Done():
		// 3. Ha a contextet törölték vagy lejárt az idő
		return "", ctx.Err()
	case res := <-ch:
		// 4. & 5. Ha kész a SubTask
		if res.err != nil {
			return "", res.err
		}
		return "Main task status:" + res.val, nil
	}
}

// SubTask simulates a long running operation
func SubTask(ctx context.Context) (string, error) {
	// Szimulálunk 200ms munkát, DE közben figyeljük a contextet is!
	select {
	case <-time.After(200 * time.Millisecond):
		// Sikeres lefutás
		return "Subtask completed successfully", nil
	case <-ctx.Done():
		// Ha időközben törölték a contextet (pl. a StartTask-ban lejárt az 1mp,
		// vagy a főprogram leállította), azonnal kilépünk.
		return "", ctx.Err()
	}
}
