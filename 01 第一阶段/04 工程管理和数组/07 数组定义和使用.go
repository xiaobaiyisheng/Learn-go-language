package main

import "fmt"

func main() {

	//var 数组名 [元素个数] 数据类型   var arr [10]int

	//var arr [10] int = [10]int {1,2,3,4,5}

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//arr:= [...]int {1,2,3,4,5,6}

	fmt.Println("数组长度：", len(arr))

	for _, i2 := range arr {
		fmt.Println(i2)
	}

}
