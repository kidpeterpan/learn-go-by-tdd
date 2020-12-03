package iteration

const repeatCount = 5

func Repeat(character string) string {
	res := ""
	for i := 0; i < repeatCount; i++ {
		res += character
	}
	return res
}
