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

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case chinese:
		prefix = chineseHelloPrefix
	}

	return prefix + name
}
