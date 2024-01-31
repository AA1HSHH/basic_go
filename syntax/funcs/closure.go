package main

func closure(name string) func() string {
	return func() string {
		return "hello," + name
	}
}

//func main() {
//	c := closure("hello")
//	println(c())
//
//}
