package readsecretregister

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// parseChannelControlRegister constructs 4 octets (8-bit long uint) based on the parameter register.
func parseChannelControlRegister(charCtrl uint32) (uint8, uint8, uint8, uint8) {
	// 1. Kinyerjük az értékeket biteltolással (shifting)
	// Byte 1 (Legfelső 8 bit): RX_PCODE
	rxPcode := uint8(charCtrl >> 24)

	// Byte 2: RX_CHAN
	rxChan := uint8(charCtrl >> 16)

	// Byte 3: TX_CHAN
	txChan := uint8(charCtrl >> 8)

	// Byte 4 (Legalsó 8 bit): TX_PCODE
	txPcode := uint8(charCtrl)

	// 2. Visszaadjuk őket a kért, összekevert sorrendben
	return txChan, rxChan, rxPcode, txPcode
}
