package main

import "fmt"

type User struct {
	Age  int
	Name string
}

func (u User) changeName(name string) {
	fmt.Printf("changeName :%p\n", &u)
	u.Name = name
}
func (u *User) changeAge(age int) {
	fmt.Printf("changeAge :%p\n", u)
	u.Age = age
}

//func main() {
//	//u1 := &User{}
//	//fmt.Println(u1)
//	//u2 := new(User)
//	//fmt.Println(u2)
//	//u3 := User{}
//	//fmt.Println(u3)
//	u4 := User{
//		Age:  18,
//		Name: "abc",
//	}
//	fmt.Printf("u :%p:%v\n", &u4, u4)
//
//	u4.changeName("123")
//	fmt.Printf("uN :%p:%v\n", &u4, u4)
//
//	u4.changeAge(10)
//	fmt.Printf("uA :%p:%v\n", &u4, u4)
//}
