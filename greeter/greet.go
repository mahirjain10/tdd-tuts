package greeter

const english = "English"
const spanish = "Spanish"
const french = "French"
const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "

func mapPrefixToLang(language string) string {
	switch language {
	case english:
		return englishPrefix
	case french:
		return frenchPrefix
	case spanish:
		return spanishPrefix
	default:
		return englishPrefix
	}
}
func Hello(name, language string) string {
	prefix := mapPrefixToLang(language)
	ifNameisEmpty := "World"
	if name == "" {
		return prefix + ifNameisEmpty
	}
	return prefix + name
}