package main

import "fmt"

const englishPrefix = "Hello, "

func main() {
	fmt.Println("hello world!") // this is side effect for test we need to return string instead

	say := Hello("Deku")
	fmt.Println(say)
}

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return englishPrefix + name
}
