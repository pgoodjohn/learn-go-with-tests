package iteration

const repeatCount = 5

func Repeat(input string) (repeated string) {

	for i := 0; i < repeatCount; i++ {
		repeated += input
	}

	return
}
