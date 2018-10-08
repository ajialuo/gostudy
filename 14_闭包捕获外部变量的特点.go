package main

import "fmt"

func main() {
	a := 10
	str := "哈哈"

	func() {
		// 闭包以引用方式捕获外部变量
		a = 666
		str = "嘿嘿嘿"
		fmt.Printf("内部:a=%d,str=%s\n", a, str)
	}()

	fmt.Printf("外部:a=%d,str=%s\n", a, str)
}
