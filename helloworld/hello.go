package main

import (
	"fmt"
	"strings"
)

const englishPrefix = "Hello, "
const thaiPrefix = "สวัสดี, "

func main() {
	fmt.Println("hello world!") // this is side effect for test we need to return string instead

	say := Hello("english", "Deku")
	fmt.Println(say)
}

func Hello(language string, name string) string {
	prefix := ""
	switch strings.ToLower(language) {
	case "english":
		prefix = englishPrefix
	case "thai":
		prefix = thaiPrefix
	default:
		prefix = englishPrefix
	}

	if name == "" {
		name = "world"
	}
	return prefix + name
}
