//声明4种结构体，并用统一的接口调用

package main

import (
	"fmt"
)

type Person interface {            //人接口
	speak(word string)
}
type Dove struct {                           //以下分别创建是四种结构体，并且为其添加方法
}
func (aDove Dove) speak(word string) {
	fmt.Println("鸽子:咕咕咕?")
}
type Repeater struct {
}
func (aRepeater Repeater) speak(word string) {
	fmt.Println("复读机:",word)
}
type Lmoner struct {
}
func (aLmoner Lmoner) speak(word string) {
	fmt.Println("柠檬精:你才是柠檬精!")
}
type Frager struct {
}
func (aFrager Frager) speak(word string) {
	fmt.Println("真香怪:真香！")
}

func main() {
	var persons = []Person{                 //分别用person接口创建四个结构体实例，并储存
		Dove{},
		Repeater{},
		Lmoner{},
		Frager{},
	}

	var word string
	fmt.Scanf("%s",&word)

	for _, person := range persons {               //分别调用四个实例的方法
        person.speak(word)
	}
}




