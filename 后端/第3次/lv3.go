package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//创建文件
	file, err0 := os.Create("./proverb.txt")
	if err0 != nil {
		fmt.Println(err0)
	}

	fmt.Println("创建文件成功!")
	defer file.Close()

    //写入文件
	content := []byte("don't communicate by sharing memory share memory by communicating")
	err1 := ioutil.WriteFile("./proverb.txt", content, 0777)
	if err1 != nil {
		fmt.Println(err1)
	}

	fmt.Println("写入文件成功!")
	defer file.Close()

    //打开并读取文件
    body, err2 := ioutil.ReadFile("proverb.txt")
    if err2 != nil {
    	fmt.Println(err2)
	}

    fmt.Println("打开并读取文件成功!")
    fmt.Println(string(body))


}
