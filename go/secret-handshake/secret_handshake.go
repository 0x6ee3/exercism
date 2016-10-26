package secret

const testVersion = 1

const (
	wink uint = 1 << iota
	doubleBlink
	closeEyes
	jump
	reverse
)

// Handshake ...
func Handshake(code uint) []string {
	var secret []string

	if code&wink == wink {
		secret = append(secret, "wink")
	}
	if code&doubleBlink == doubleBlink {
		secret = append(secret, "double blink")
	}
	if code&closeEyes == closeEyes {
		secret = append(secret, "close your eyes")
	}
	if code&jump == jump {
		secret = append(secret, "jump")
	}
	if code&reverse == reverse {
		lenSecret := len(secret)
		for i, j := 0, 0; i < lenSecret/2; i++ {
			j = lenSecret - 1 - i
			secret[i], secret[j] = secret[j], secret[i]
		}
	}

	return secret
}
