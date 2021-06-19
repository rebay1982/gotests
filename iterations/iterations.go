package iterations

const repeatCount = 5

// Repeat a repeat function man.
func Repeat(character string) string {
	var repeated string

	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated
}
