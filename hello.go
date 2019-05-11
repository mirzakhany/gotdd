package main

import "fmt"

// HelloPrefix is hello prefix
const HelloPrefix = "Hello, "

// Hello return a hello message
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return HelloPrefix + name
}

func main() {

	fmt.Println(Hello("mohsen"))
}
