package iteration

const repeatCount = 5

// Repeat takes in a string and returns that string repeated 5 times
func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
