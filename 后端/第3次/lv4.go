package main

import (
	"fmt"
)

var myRes = make(chan int)

func factorial(n int) {
	var res = 1
	for i := 1; i <= n; i ++ {
		res *= i
	}
	myRes <- res
}

func Endfactorial(n int) {
	var res = 1
	for i := 1; i <= n; i ++ {
		res *= i
	}
	myRes <- res
	close(myRes)
}

func main() {
	for i := 1; i <= 20; i ++ {              //前19次阶乘使用factorial函数,没有关闭myRes,如果关闭,则下一次传入时会出错,而如果
		if i != 20 {                         //不关闭myRes,在遍历的时候会出现错误,如果在main协程中关闭myRes,main协程的速度比较快
		    go factorial(i)                  //算阶乘的协程还没有来得及计算就关闭了myRes,同样会出错,因此只好在最后一次计算阶乘时
		} else {                             //为保证计算完毕,故在计算函数结束后关闭myRes,因此把最后一次单独声明了一个函数
			go Endfactorial(i)
		}

	}

	for v := range myRes {                  //为保证遍历时不出错,遍历之前必须关闭管道,但是也得计算完成,因此引入管道来存储数据
		fmt.Println(v)                      //因为在遍历时,若管道没有写入结束,是处于堵塞状态的,也就不会导致主线程结束都分快,而
	}                                       //factorial协程没有结束的情况了
}


