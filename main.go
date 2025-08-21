package main

const englishHelloPrefix = "Hello, "
const chineseHelloPrefix = "你好， "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Chinese" {
		return chineseHelloPrefix + name
	}

	return englishHelloPrefix + name
}
