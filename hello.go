package main

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return helloPrefix(language) + name
}

const spanish = "Spanish"
const helloPrefixSpanish = "Hola, "
const helloPrefixEnglish = "Hello, "

func helloPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = helloPrefixSpanish
	default:
		prefix = helloPrefixEnglish
	}

	return
}
