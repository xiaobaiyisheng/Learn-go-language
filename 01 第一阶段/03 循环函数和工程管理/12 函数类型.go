package main

import "fmt"

// type 可以定义函数类型
// type 可以为已存在的函数起别名

func test1201() {
	fmt.Println("不带参数的test1201()")
}

func test1202(a, b int) {
	fmt.Printf("调用函数 test1202(a,b int) \n 当前传递的参数是 test1202(%d,%d)\n", a, b)
}

// 起个别名
type FUNCtest func()

//type FUNCtest1 func(a,b int)

func main() {
	//定义函数类型变量
	var f1 FUNCtest
	f1 = test1201
	f1()

	f2 := test1202
	f2(1, 2)

	//fmt.Println(test1203(1, 2, 3))
}
