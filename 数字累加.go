package main

import "fmt"

func test01() (sum int) {
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return
}

// 通过递归实现
func test02(i int) int {
	if i == 1 {
		return 1
	}
	return i + test02(i-1)
}

func main() {
	var sum int
	//sum = test01()
	sum = test02(100)
	fmt.Println("sum=", sum)
}
