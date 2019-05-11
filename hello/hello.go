package hello

// Consts
const (
	Spanish            = "Spanish"
	English            = "English"
	Franch             = "Franch"
	SpanishHelloPrefix = "Hola, "
	FranchHelloPrefix  = "Hi, "
	EnglishHelloPrefix = "Hello, "
)

// Hello return a hello message
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greatingPrefix(language) + name
}

func greatingPrefix(language string) (prefix string) {
	switch language {
	case Franch:
		prefix = FranchHelloPrefix
	case Spanish:
		prefix = SpanishHelloPrefix
	default:
		prefix = EnglishHelloPrefix
	}
	return
}
