package channelbroadcaster

import "context"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE

// channelBroadcast broadcasts input values to multiple output channels using non-blocking sends.
func channelBroadcast(ctx context.Context, input <-chan any, outputs []chan<- any) {
	// 1. Indítunk egy saját goroutine-t a feldolgozáshoz
	go func() {
		// 4. Biztosítjuk, hogy kilépéskor minden kimeneti csatorna lezárásra kerüljön
		defer func() {
			for _, ch := range outputs {
				close(ch)
			}
		}()

		for {
			select {
			// 3. Context cancellation figyelése
			case <-ctx.Done():
				return // Kilépünk, a defer lezárja a kimeneteket

			// 2. & 4. Olvasás a bemeneti csatornából
			case val, ok := <-input:
				if !ok {
					return // Ha a bemenetet lezárták, mi is leállunk
				}

				// 2. & 5. Szétszórás (Broadcast) minden kimenetre
				for _, out := range outputs {
					select {
					case out <- val:
						// Sikeres küldés
					default:
						// 5. szabály: "acceptable to potentially miss sends".
						// Ha a kimenet tele van (blokkolna), akkor átugorjuk ezt a csatornát,
						// hogy ne tartsuk fel a többit.
					}
				}
			}
		}
	}()
}
