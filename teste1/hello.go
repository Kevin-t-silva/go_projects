package main

import "fmt"

const englishHelloprefix = "Hello, "
const chooseLanguage = "Portuguese"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == chooseLanguage {
		return "Olá, " + name
	}
	return englishHelloprefix + name
}
func main() {
	fmt.Println(Hello("", ""))
}
