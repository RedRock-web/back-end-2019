//利用心型函数打印心形
package main

import "fmt"
var x,y float64                      //以x，y为横纵坐标

func main()  {
	for y := -1.5; y < 1.5; y = y + 0.14 {
		for x := -1.5; x < 1.5; x = x + 0.05 {
			sum := (x * x + y * y - 1) * (x * x + y * y - 1) * (x * x + y * y - 1) + x * x * y * y * y
			if sum <= 0.0 {                                          //若满足点(x,y)在心型函数内，则输出，否则输出空格
				fmt.Printf("I")
			}else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}

}


