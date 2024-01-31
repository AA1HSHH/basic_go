package main

import "fmt"

func Array() {
	a1 := [3]int{1, 2, 3}
	fmt.Printf("a1:%v, len: %d, cap: %d", a1, len(a1), cap(a1))
}

//func main() {
//	Array()
//}
