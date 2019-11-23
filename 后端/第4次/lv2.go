//输出时间戳,然后返回具体时间
package main

import (
	"fmt"
	"time"
)

func main() {
	TimeStampTransform()
}

func TimeStampTransform() {
	var timeStampInput int64
	var timeArr []int64

	fmt.Println("请输入时间戳,将返回具体时间(输入result结束并输入结果)")
	for {
		judge, _ := fmt.Scanf("%d", &timeStampInput)
		//fmt.Println(judge)
		if judge == 0 {
			fmt.Println("the result are :")
			for j := 0; j < len(timeArr); j ++ {
				fmt.Println(time.Unix(timeArr[0], 0))
			}
			break
		} else {
			timeArr = append(timeArr, timeStampInput)
			fmt.Println("input ok!")
		}
	}
}
