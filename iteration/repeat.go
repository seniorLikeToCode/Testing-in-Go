package iteration

func Repeat(count int, charachter string) string {
	var repeated string
	for range count {
		repeated += charachter
	}

	return repeated
}
