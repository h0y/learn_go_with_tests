package main

import "fmt"

const english = "English"
const chinese = "Chinese"

const englishPrefix = "Hello "
const chinesePrefix = "你好 "
const symbolSuffix = "!"

func Hello(name, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name + symbolSuffix
}

func greetingPrefix(language string) (prefix string) {
	prefix = englishPrefix

	switch language {
	case english:
		prefix = englishPrefix
	case chinese:
		prefix = chinesePrefix
	}

	return prefix
}

func main() {
	fmt.Println(Hello("Ellery", "Chinese"));
}