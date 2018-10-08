package main

import "fmt"

func main() {
	a := 10
	str := "jack"

	// 匿名函数，没有函数名字，函数定义
	f1 := func() {
		fmt.Println("a=", a)
		fmt.Println("str=", str)
	}

	f1()

	// 定义匿名函数，同时调用
	func() {
		fmt.Printf("a= %d,str= %s\n", a, str)
	}()

	// 带参数的匿名函数
	f3 := func(i, j int) {
		fmt.Printf("i=%d,j=%d\n", i, j)
	}
	f3(1, 2)

	// 带参数的匿名函数,同时调用
	func(i, j int) {
		fmt.Printf("i=%d,j=%d\n", i, j)
	}(3, 4)

	// 匿名函数，有参有返回值
	x, y := func(i, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return
	}(10, 20)
	fmt.Printf("x=%d,y=%d\n", x, y)
}
