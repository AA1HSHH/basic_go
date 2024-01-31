package main

func main() {
	var a byte = 'a'
	println(a)

	var str string = "hello"
	var bs []byte = []byte(str)
	var str1 string = string(bs)
	println(str1)
}
