package main

import "fmt"
func main() {
	var (
		i int = 0;
		k int = 0;
		j int = 0;
		num int = 0;)
	for i = 0; i < 4; i ++ {
		for j = 0; j < 4; j++ {
			for k = 0; k < 4; k++ {
				if i != j && i != k && j != k {
					num++
					fmt.Println(i, j, k)
				}
			}
		}
	}
	fmt.Printf("一共有%d种",num)
}




