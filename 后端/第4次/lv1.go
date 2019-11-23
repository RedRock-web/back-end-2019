//学院-班级-学生之间的相互查询
//学院:计算机学院,1班张三学号10000,李四10001,二班周星10003,陈丽10004
//学院:法律学院,3班文心10005,石书10006,4班齐易10007,毛志10008
package main

import "fmt"

func main() {
	Universal()
}

func Universal() {
	var (
		choice      int
		studentName string
		studentId   int
		classNum    int
	)

	studentNameToClass := map[string]int{
		"张三": 1,
		"李四": 1,
		"周星": 2,
		"陈丽": 2,
		"文心": 3,
		"石书": 3,
		"齐易": 4,
		"毛志": 4,
	}
	studentNameToCollege := map[string]string{
		"张三": "计算机学院",
		"李四": "计算机学院",
		"周星": "计算机学院",
		"陈丽": "计算机学院",
		"文心": "法律学院",
		"石书": "法律学院",
		"齐易": "法律学院",
		"毛志": "法律学院",
	}
	studentIdToClass := map[int]int{
		19001: 1,
		19002: 1,
		19003: 2,
		19004: 2,
		19005: 3,
		19006: 3,
		19007: 4,
		19008: 4,
	}
	studentIdToCollege := map[int]string{
		19001: "计算机学院",
		19002: "计算机学院",
		19003: "计算机学院",
		19004: "计算机学院",
		19005: "法律学院",
		19006: "法律学院",
		19007: "法律学院",
		19008: "法律学院",
	}

	classAll := map[int]map[string]int{
		1: {
			"张三": 19001,
			"李四": 19002,
		},
		2: {
			"周星": 19003,
			"陈丽": 19004,
		},
		3: {
			"文心": 19005,
			"石书": 19006,
		},
		4: {
			"齐易": 19007,
			"毛志": 19008,
		},
	}

	fmt.Println("1输入姓名查询信息\n2输入学号查询信息\n3查询班级信息")
rego:
	fmt.Scanf("%d", &choice)

	if choice == 1 {
		fmt.Println("请输入你的名字:")
		fmt.Scanf("%s", &studentName)
		fmt.Printf("%s,你的学院为%s,班级为%d班", studentName, studentNameToCollege[studentName], studentNameToClass[studentName])
	} else if choice == 2 {
		fmt.Println("请输入你的学号:")
		fmt.Scanf("%d", &studentId)
		fmt.Printf("%d,你的学院为%s,班级为%d班", studentId, studentIdToCollege[studentId], studentIdToClass[studentId])

	} else if choice == 3 {
		fmt.Printf("请输入班级号:")
		fmt.Scanf("%d", &classNum)
		if classNum >= 1 && classNum <= 4 {
			fmt.Println(classAll[classNum])
		} else {
			fmt.Println("没有此班级!")
		}
	} else {
		fmt.Println("输入不在范围内,请重试")
		goto rego
	}
}
