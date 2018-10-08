package main

import "fmt"

type FuncType func(int, int) int

// 实现加法
func Add(a, b int) int {
	return a + b
}

// 实现减法
func Minus(a, b int) int {
	return a - b
}

func Calc(a, b int, fTest FuncType) (result int) {
	fmt.Println("calc")
	result = fTest(a, b)
	return
}

func main() {
	a := Calc(1, 1, Add)
	fmt.Println("a=", a)
	b := Calc(1, 1, Minus)
	fmt.Println("b=", b)
}
