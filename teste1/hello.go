package main

import "fmt"

const englishHelloprefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloprefix + name
}
func main() {
	fmt.Println(Hello(""))
}
