package iteration

func Repeat(input string, repeatCount int) (repeated string) {

	for i := 0; i < repeatCount; i++ {
		repeated += input
	}

	return
}
