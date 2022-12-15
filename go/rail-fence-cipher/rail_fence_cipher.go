package railfence

// Encode uses rail fence cipher to output encrypted text based on message and nuber of rails.
func Encode(message string, rails int) string {
	var result string
	inc := 1
	osc := 0
	a := 0
	b := rails - 1
	fence := make([]string, rails)

	// Put message into zigzag shape of the fence.
	for _, v := range message {
		fence[osc] += string(v)
		osc, inc = oscillate(osc, inc, a, b)
	}

	for _, v := range fence {
		result += v
	}
	return result
}

// Decode uses rail fence cipher to output clear text based on encrypted message and nuber of rails.
func Decode(message string, rails int) string {
	var result string
	inc := 1
	osc := 0
	a := 0
	b := rails - 1
	l := len(message)
	fence := make([][]string, rails)

	// Make empty fence.
	for i := range fence {
		fence[i] = make([]string, l)
	}

	// Put spaces onto fence in the zigzag shape.
	for j := range fence[0] {
		fence[osc][j] = " "
		osc, inc = oscillate(osc, inc, a, b)
	}

	// Put chars of encrypted text onto the fence - rows first.
	for i := range fence {
		for j := range fence[i] {
			if fence[i][j] == " " {
				fence[i][j] = string(message[0])
				message = message[1:]
			}
		}
	}

	osc = 0

	// Collect clear text message from zigzag shape.
	for j := range fence[0] {
		result += fence[osc][j]
		osc, inc = oscillate(osc, inc, a, b)
	}
	return result
}

func oscillate(osc, inc, a, b int) (int, int) {
	if osc == a {
		inc = 1
	} else if osc == b {
		inc = -1
	}
	osc += inc
	return osc, inc
}
