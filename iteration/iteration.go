package iteration

// Repeat repeats the str count times
func Repeat(count int, str string) string {
	var ret string
	for i := 0; i < count; i++ {
		ret += str
	}
	return ret
}
