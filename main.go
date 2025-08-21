package main

const englishHelloPrefix = "Hello, "
const chineseHelloPrefix = "你好， "
const spanishHelloPrefix = "Hola, "

const spanish = "Spanish"
const chinese = "Chinese"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == chinese {
		return chineseHelloPrefix + name
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}

	return englishHelloPrefix + name
}
