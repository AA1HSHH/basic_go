package main

func F3(abc string) {
	println("hello", abc)
}

func F6() func(name string) string {
	return func(name string) string {
		return "hello" + name
	}
}
func F7() {
	fn := func(name string) string {
		return "hello " + name
	}("abcv3")
	println(fn)
}

//func main() {
//	f3 := F3
//	f3("abc")
//
//	fn := func(name string) string {
//		return "hello " + name
//	}
//
//	str := fn("abc1")
//	println(str)
//
//	f6 := F6()
//	println(f6("abc2"))
//
//	F7()
//}
