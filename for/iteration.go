package iteration

const repeatCount = 5

func repeat(s string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += s
	}

	return repeated
}

func exampleRepeat(s string, t int) string {
	var repeated string
	for i := 0; i < t; i++ {
		repeated += s
	}
	return repeated
}
