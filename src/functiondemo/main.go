package main

import (
	"fmt"
	"functiondemo/SumAndSub"
)

func Sum(n1 int, n2 int) int {
	sum := n1 + n2
	return sum

}
func main() {
	fmt.Println("你好")

	fmt.Println("res=", Sum(10, 20))

	fmt.Println("res2=", SumAndSub.Sub(20, 30))

}
