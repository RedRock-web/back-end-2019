//输入字符串，输出其倒序
package main
import "fmt"

func main() {
	var (
		voc string
		arr []string
		arr_sort []string
		)

	fmt.Scanf("%s",&voc)               //输入字符串并储存在变量中
	for i := 0; i < len(voc); i ++ {               //将字符串元素储存在切片中

		arr = append(arr, string(voc[i]))
	}
/*	fmt.Println(arr)
	fmt.Println(len(arr))                //检查数组长度
*/

	for c := 0; c < len(arr); c ++ {
		arr_sort = append(arr_sort,arr[len(arr) - c -1])       //将切片倒序

	}
	//fmt.Println(arr_sort)
	for a := 0; a < len(arr_sort); a ++ {
		fmt.Printf("%s",arr_sort[a])                //倒序输出
	}




























}
