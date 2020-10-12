package main

import "fmt"

func main() {
	a := 10
	b := 20

	//匿名内部函数
	//func(a,b int){
	//	fmt.Println(a+b)
	//}(a,b)

	// 匿名内部函数的拆解
	//把函数赋值给f，f = 函数变量
	f := func(a, b int) {
		fmt.Println(a + b)
	}

	f(a, b)

	// 有返回值的匿名函数

	fv := func(a, b int) int {
		return a + b
	}

	fv1 := fv(1, 2)

	fmt.Println(fv1)

	//无参，无反，直接调用的匿名函数写法 1
	{

	}
	//无参，无返，直接调用的匿名函数写法 2
	func() {

	}()

}
