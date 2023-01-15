package go_say_hello

func SayHello() string {
	return "Hello"
}

func SayHelloTo(name string) string {
	greet := "Hello " + name
	return greet
}
