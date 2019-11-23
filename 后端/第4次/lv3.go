//找出学校里重名最多的学生
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

//var wg sync.WaitGroup

func main() {
	nameArry, nameMap := GetNameList()
	SaveData("studentIdToName.txt", nameMap)
	MostName(nameArry)
}

//传入数据,将其转换为json后再转换为string后存入文件中
func SaveData(fileName string, saveData interface{}) {
	saveDataJson, _ := json.Marshal(saveData)
	dataString := string(saveDataJson)

	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
	}
	file.WriteString(dataString)
}

//读取文件,然后将其转换为原本的格式
func ReadData(dataName string) interface{} {
	data, err := ioutil.ReadFile(dataName)
	if err != nil {
	}
	fmt.Println("string类型:")
	fmt.Println(string(data))
	fmt.Println("是否进一步转换?\n1.map[string]string\n2.map[string]int\n3.map[int]int\n4.map[int]string")
	var choice int
	var returnData interface{}

	fmt.Scanf("%d", &choice)
	switch choice {
	case 1:
		data1 := make(map[string]string)
		err := json.Unmarshal([]byte(data), &data1)
		if err != nil {
		}
		//case 1:	   待扩展
		//fmt.Println(data1)
	}

	return returnData
}

func GetName(id int64) string {
	defer func() {
		recover()
	}()
	return ChineseTrans(SplitName(Grab(id)))
}

func GetNameList() ([]string, map[int64]string) {
/*	defer func() {
		recover()
	}()*/
	var (
		nameListTotal  []string
		nameListTotal1 []string
		nameListTotal2 []string
		nameList1      []string
		nameList2      []string
		nameList3      []string
		nameList4      []string
		id1            int64
		id2            int64
		id3            int64
		id4            int64
		lock sync.Mutex
	)
	IdToName :=make(map[int64]string)

	ch := make(chan struct{})
	count := 4

	//wg.Add(1)
	go func() {
		//wg.Add(1)
		defer func() {
			recover()
		}()
		for id1 = 2016210001; id1 < 2016216001; id1 ++ {
			aName := GetName(id1)
			if aName != "" {
				nameList1 = append(nameList1, aName)
				lock.Lock()
				IdToName[id1] = aName
				lock.Unlock()
				fmt.Println(id1, aName)
			}
		}

		ch <- struct{}{}
		//defer wg.Add(-1)
	}()

	go func() {
		//wg.Add(1)
		defer func() {
			recover()
		}()
		for id2 = 2017210001; id2 < 2017216001; id2 ++ {
			bName := GetName(id2)
			if bName != "" {
				nameList2 = append(nameList2, bName)
				lock.Lock()
				IdToName[id2] = bName
				lock.Unlock()
				fmt.Println(id2, bName)
			}
		}
		ch <- struct{}{}
		//defer wg.Add(-1)
	}()

	go func() {
		defer func() {
			recover()
		}()
		//wg.Add(1)
		for id3 = 2018210001; id3 < 2018216001; id3 ++ {
			cName := GetName(id3)
			if cName != "" {
				nameList3 = append(nameList3, cName)
				lock.Lock()
				IdToName[id3] = cName
				lock.Unlock()
				fmt.Println(id3, cName)
			}
		}
		ch <- struct{}{}
		//defer wg.Add(-1)
	}()

	go func() {
		defer func() {
			recover()
		}()
		//wg.Add(1)
		for id4 = 2019210001; id4 < 2019216001; id4 ++ {
			dName := GetName(id4)
			if dName != "" {
				nameList4 = append(nameList4, dName)
				lock.Lock()
				IdToName[id4] = dName
				lock.Unlock()
				fmt.Println(id4, dName)
			}
		}
		ch <- struct{}{}
		//defer wg.Add(-1)
	}()

	for range ch {
		count --
		if count == 0 {
			close(ch)
		}
	}

	//fmt.Println(nameList)
	nameListTotal1 = append(nameList1, nameList2...)
	nameListTotal2 = append(nameList3, nameList4...)
	nameListTotal = append(nameListTotal1, nameListTotal2...)

	//fmt.Println(nameListTotal)
	//defer wg.Add(-1)
	//wg.Wait()

	return nameListTotal, IdToName
}

//var count int

//找到数组中重复最多的元素
func MostName(nameList []string) (string, int) {
	/*	defer func() {
			recover()
		}()
	*/
	repeatNumList := make(map[string]int)
	//count := 0
	//listLen := len(nameList)

	for i := 0; i < len(nameList); i ++ {
		//fmt.Println(nameList)
		name := nameList[i]
		//fmt.Println()
		//fmt.Printf("i = %d,name = %s, count = %d\n", i, name, count)

		for j := i + 1; j < len(nameList); j ++ {
			nameNext := nameList[j]

			if name == nameNext {
				repeatNumList[name] ++
				//nameList = append(nameList[:j], nameList[j+1:]...)
			}
			//fmt.Println(nameList)
			//fmt.Printf("j = %dnameNext = %s, count = %d\n", j, nameNext, count)
		}
		//nameList = append(nameList[:i], nameList[i+1:]...)
	}

	for key, value := range repeatNumList {

		for n := 1; ; n ++ {
			if (n*(n-1))/2 == value {
				repeatNumList[key] = n
				break
			}
		}
	}

	//fmt.Println(repeatNumList)
	//var maxNumKey string
	maxNumValue := 0

	for _, v2 := range repeatNumList {
		//fmt.Printf("k1= %s,v1=%d, v2= %d\n", k1, v1, v2)
		if maxNumValue < v2 {
			maxNumValue = v2

		}
	}
	theMostName := ValueFoundKey(maxNumValue, repeatNumList)

	if maxNumValue == 0 {
		fmt.Println("没有重复的名字")
	} else {
		fmt.Printf("名字重复最多的是%s, 一共有%d个\n", theMostName, maxNumValue)
	}
	return theMostName, maxNumValue
}

func ValueFoundKey(value interface{}, targeMap map[string]int) string { //通过map的value找到对应的第一个匹配的key值
	var targeKey string
	for k, v := range targeMap {
		if value == v {
			targeKey = k
			break
		}
	}
	return targeKey
}

func Grab(id int64) string { //get传入学号,获取json返回
	resp, err := http.Get("http://jwzx.cqupt.edu.cn/data/json_StudentSearch.php?searchKey=" +
		strconv.FormatInt(id, 10))
	if err != nil {
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		panic("error")
	}

	return string(body)
}

func SplitName(s string) string { //从得到的json截取xm片段
	defer func() {
		recover()
	}()
	tempSlice := strings.FieldsFunc(string(s), SplitOne)

	name := strings.FieldsFunc(tempSlice[3], SplitTwo)
	return string(name[2])
}

func SplitOne(s rune) bool {
	if s == ',' {
		return true
	}
	return false
}

func SplitTwo(s rune) bool {
	if s == '"' {
		return true
	}
	return false
}

func ChineseTrans(word string) string { //将中文的Unicode字符串转换为中文,主要是有个转义字符\要消掉
	var context string
	defer func() {
		recover()
	}()

	sUnicodev := strings.Split(word, "\\u")
	for _, v := range sUnicodev {
		if len(v) < 1 {
			continue
		}
		temp, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			panic(err)
		}
		context += fmt.Sprintf("%c", temp)
	}
	return context
}
