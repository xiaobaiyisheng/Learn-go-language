package main

import "fmt"

// 函数在调用结束会从内存中销毁
func test1(a int) {
	a++
	fmt.Println(a)
}

func main0301() {

	a := 1
	for i := 0; i < 10; i++ {
		test1(a)
	}

}

//可以通过匿名函数和闭包，实现函数在战区的本地化
func test2(a int) func() int {

	return func() int {
		a++
		return a
	}

}

func main() {
	var v1 = 0

	// 该变量的释放周期 = main() 的生存周期
	f := test2(v1)

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

}
