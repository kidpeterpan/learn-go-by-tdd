package main

import "fmt"

func main() {
	fmt.Println("hello world!") // this is side effect for test we need to return string instead
}

func Hello() string {
	return "hello world!"
}