package main

func Forslice() {
	a := []int{1, 2, 3}
	for i, val := range a {
		println(i, val)
	}
}

type User struct {
	name string
}

func LoopBug() {
	users := []User{
		{
			name: "abc",
		},
		{
			"123",
		},
	}
	m := make(map[string]*User)
	// u的地址是同一个
	for _, u := range users {
		m[u.name] = &u
	}
	for k, v := range m {
		println(k, v.name)
	}
}
func main() {
	LoopBug()
}
