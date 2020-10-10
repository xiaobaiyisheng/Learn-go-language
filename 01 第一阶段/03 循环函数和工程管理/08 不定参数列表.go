package main

import (
	"fmt"
)

// 不定参函数，函数类型为 ...数据类型
func test(args ...int) {
	fmt.Println("传入的参数的内容是：", args)

	// 借助 len() 计算不定参函数的长度
	for i := 0; i < len(args); i++ {
		fmt.Printf("下标 %d 的值是：%d \n", i, args[i])
	}
}

// 加法计算模型
func plus(arr ...int) int {

	sum := 0
	fmt.Println("求下列整数类型数字的累加结果：", arr)

	//for i := 0; i < len(arr); i++ {
	//	sum += arr[i]
	//}

	// range 可以遍历 集合种的数据信息，数组，切片，结构体数组，map
	for _, i := range arr {
		sum += i
	}
	return sum

}

func main() {

	test(1, 2)
	//test(1,2,3,4)
	//test(1,2,3,4,true)  // 08 不定参数列表.go:14:13: cannot use true (type bool) as type int in argument to test

	// 累加计算（将累加结果直接打印了，没有用变量去接收）
	fmt.Println(plus(1, 2, 3, 4))

}
