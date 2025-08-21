package main

const englishHelloPrefix = "Hello, "
const chineseHelloPrefix = "你好， "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Chinese" {
		return chineseHelloPrefix + name
	}

	if language == "Spanish" {
		return spanishHelloPrefix + name
	}

	return englishHelloPrefix + name
}
