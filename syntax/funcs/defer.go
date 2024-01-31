package main

func DeferClosure() {
	i := 0
	defer func() {
		println(i)
	}()
	i++
}
func DeferClosurev1() {
	i := 0
	defer func(i int) {
		println(i)
	}(i)
	i++
}

func Deferchange() int {
	a := 0
	defer func() {
		a = 1
	}()
	return a
}
func DeferchangeV1() (a int) {
	a = 0
	defer func() {
		a = 1
	}()
	return a
}

type Mystruct struct {
	name string
}

func DeferchangeV2() *Mystruct {
	res := &Mystruct{
		name: "Tom",
	}
	defer func() {
		res.name = "ABC"
	}()
	return res
}
func DeferLoopV1() {
	for i := 0; i < 10; i++ {
		defer func() {
			println(i)
		}()
	}
}
func DeferLoopV2() {
	for i := 0; i < 10; i++ {
		defer func(i int) {
			println(i)
		}(i)
	}
}
func DeferLoopV3() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			println(j)
		}()
	}
}
func test() {
	j := 0
	for i := 0; i < 10; i++ {
		j = i
		defer func() {
			println(j)
		}()
	}
}
func main() {
	//DeferClosure()
	//println(".....")
	//DeferClosurev1()
	////println(Deferchange())
	////println(DeferchangeV1())
	////println(DeferchangeV2().name)
	//println(".....")
	//DeferLoopV1()
	//println(".....")
	//DeferLoopV2()
	//println(".....")
	//DeferLoopV3()
	test()

}
