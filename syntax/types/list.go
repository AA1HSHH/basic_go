package main

type List interface {
	Add(idx int, val any)
	Append(val any)
	Delete(idx int)
}

type LinkedList struct {
	Head node
}

func (l *LinkedList) Append(val any) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) Delete(idx int) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) Add(idx int, val any) {

}

type node struct {
}

func main() {
	l := &LinkedList{}
	l.Add(1, 123)
	l.Add(1, "123")
}
