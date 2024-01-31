package main

import "fmt"

func Slice() {
	//s1 := []int{1, 2, 3}
	s2 := make([]int, 3, 4)
	for i := 0; i < 50; i++ {
		s2 = append(s2, i)
		fmt.Printf("s2 len:%d, cap:%d\n", len(s2), cap(s2))
	}
	//fmt.Printf("s1:%v len:%d, cap:%d\n", s1, len(s1), cap(s1))

}
func SubSlice() {
	s1 := []int{2, 4, 6, 8, 10}
	fmt.Printf("s1:%v len:%d, cap:%d\n", s1, len(s1), cap(s1))
	s2 := s1[1:3]
	fmt.Printf("s2: %v, len:%d, cap:%d\n", s2, len(s2), cap(s2))
}
func shareSlice() {
	s1 := []int{1, 2, 3, 4}
	s2 := s1[2:]
	fmt.Printf("s1:%v len:%d, cap:%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len:%d, cap:%d\n", s2, len(s2), cap(s2))
	s2[0] = 199
	fmt.Printf("s1:%v len:%d, cap:%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len:%d, cap:%d\n", s2, len(s2), cap(s2))
	s2 = append(s2, 200)
	fmt.Printf("s1:%v len:%d, cap:%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len:%d, cap:%d\n", s2, len(s2), cap(s2))

	s2[0] = 1999
	fmt.Printf("s1:%v len:%d, cap:%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len:%d, cap:%d\n", s2, len(s2), cap(s2))
}
func main() {
	shareSlice()
}
