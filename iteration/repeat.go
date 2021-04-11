package iteration

// repeats a character reps times
func Repeat(value string, reps int) (result string) {
	for i := 0; i < reps; i++ {
		result += value
	}
	return
}
