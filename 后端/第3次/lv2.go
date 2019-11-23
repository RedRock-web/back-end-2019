//使用goroutine求1~10000以内的素数

package main

import (
	"fmt"
	"time"
)

func isPrimeNumber(num int) bool {           //判断是否一个数是否是素数
	DivisibleNum := 0

	for i := 2; i < num; i ++ {
		if num % i == 0 {
			DivisibleNum ++
		}
	}
	if DivisibleNum == 0 {
		return true
	} else {
		return false
	}

}

func PrimerNumberList(n int) []int {                 //打印出素数切片
	var primeNumberSlice []int

	for j := 1; j <= n; j ++ {
		if isPrimeNumber(j) {
			primeNumberSlice = append(primeNumberSlice, j)
		}
	}
	fmt.Println(primeNumberSlice)
	return primeNumberSlice

}



func main() {
	go PrimerNumberList(10000)

	time.Sleep(time.Second)

}


