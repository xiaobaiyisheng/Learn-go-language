package main

import (
	"fmt"
)

// 调用自己的函数，就是递归函数, 递归比较消耗资源

//死递归
func test4(a int) {

	// 递归要有明确的退出/终止标记 return
	if a == 0 {
		return
	}
	fmt.Println(a)
	a--

	test4(a)
}

func main0401() {
	test4(4)

}

var sum = 1

func test401(a int) {

	if a == 1 {
		return
	}

	test401(a - 1)
	sum *= a

}

// 计算一个数字的阶乘
func main() {

	test401(3)
	fmt.Println(sum)

	// 循环计算阶乘
	//tmp :=1
	//for i := 1; i <= 4; i++ {
	//	tmp*=i
	//}
	//fmt.Println(tmp)

}
