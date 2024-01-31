package main

import (
	"fmt"
	"io"
)

func sum[T Number](vals []T) T {
	var res T
	for _, v := range vals {
		res += v
	}
	return res
}

type Number interface {
	~int
}
type Integer int

func close[T io.Closer]() {
	var t T
	t.Close()
}
func main() {
	v1 := sum([]int{123, 321})
	fmt.Println(v1)

	v2 := sum([]Integer{123, 321})
	fmt.Println(v2)
}
