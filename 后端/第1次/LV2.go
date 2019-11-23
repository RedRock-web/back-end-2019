//利用心型函数打印心形
package main

import "fmt"
var x,y int

func main()  {
	for y := 0; y < 100; y ++ {
		for x := 0; x < 100; x ++ {
			sum := (x * x + y * y - 1) * (x * x + y * y - 1) * (x * x + y * y - 1) + x * x * y * y * y
			if sum <= 0.0 {
				fmt.Printf("I")
			}else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}

}


