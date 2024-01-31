package main

type ListV1[T any] interface {
	Add(index int, val T)
	Append(val T)
	Delete(index int)
}
type LinkedListV1[T any] struct {
	Node[T]
}

func (l LinkedListV1[T]) Add(index int, val T) {
	//TODO implement me
	panic("implement me")
}

func (l LinkedListV1[T]) Append(val T) {
	//TODO implement me
	panic("implement me")
}

func (l LinkedListV1[T]) Delete(index int) {
	//TODO implement me
	panic("implement me")
}

type Node[T any] struct {
}

func main() {
	l := &LinkedListV1[int]{}
	l.Add(1, 123)
	//l.Add(1,"123")

}
