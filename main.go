package main

import "fmt"

func main() {
	fmt.Println(Hello("Asher"))
}

func Hello(name string) string {
	return "Hello, " + name
}
