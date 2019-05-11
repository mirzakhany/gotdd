package main

import "fmt"

const HelloPrefix = "Hello, "


func Hello(name string) string {
	return HelloPrefix + name
}

func main() {

	fmt.Println(Hello("mohsen"))

}
