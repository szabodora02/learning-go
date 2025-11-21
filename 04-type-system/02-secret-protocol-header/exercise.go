package secretprotocolheader

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// createPublishFixHeader constructs an octet (8-bit long byte) based on its three arguments and the fix QoS setting.
func createPublishFixHeader(isFirstAttempt bool, isBroadcasted bool, isSecure bool) byte {
	// Kezdőérték a fix bitekből:
	// - Packet Type (010... -> 6. bit 1) = 64 (0x40)
	// - QoS (At least once -> ...01.. -> 2. bit 1) = 4 (0x04)
	// 0x40 | 0x04 = 0x44
	var h byte = 0x44

	// Bit 4: FirstAttempt
	if isFirstAttempt {
		h = h | (1 << 4)
	}

	// Bit 1: Broadcast
	if isBroadcasted {
		h = h | (1 << 1)
	}

	// Bit 0: Secure
	if isSecure {
		h = h | 1
	}

	return h
}
