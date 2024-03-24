package service

import "fmt"

type (
	IAnh interface {
		PrintHelloWorld()
	}
	Anh struct{}
)

func NewAnh() IAnh {
	return &Anh{}
}

func (anh *Anh) PrintHelloWorld() {
	fmt.Println("Hello World")
}
