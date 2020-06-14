package goheroku

// HelloWorld purpose is to provide initial gowebservice implementation
func HelloWorld() string {
	return Hello("World")
}

// Hello with parameter
func Hello(name string) string {
	return "Hello " + name
}
