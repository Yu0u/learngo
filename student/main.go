package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
)

type Student struct {
	id       int     // 学号
	name     string  // 姓名
	chinese  float64 // 语文成绩
	math     float64 // 数学成绩
	english  float64 // 英语成绩
	avgScore float64 // 平均成绩
	sumScore float64 // 总成绩
}

func NewStudent(id int, name string, chinese float64, math float64, english float64) *Student {
	return &Student{id: id, name: name, chinese: chinese, math: math, english: english}
}

func main() {
	var s []Student
	home(&s)
}

func home(s *[]Student) {
	for {
		printMenu()
		switch inputInt() {
		case 1:
			showTranscript(*s)
			fmt.Print("按Enter返回主菜单")
			fmt.Scanln()
		case 2:
			inputTranscript(s)
		case 3:
			modifyStudent(s)
		case 4:
			result := removeStudent(s)
			s = &result
		case 5:
			filterStudent(s)
		case 6:
			sortStudent(*s)
		case 7:
			os.Exit(0)
		}
	}
}

func sortStudent(students []Student) {
	sort.Slice(students, func(i, j int) bool {
		return students[i].sumScore > students[j].sumScore
	})
	fmt.Println("排序成功")
}

func removeStudent(s *[]Student) []Student {
	fmt.Println("请输入学号:")
	studentId := inputInt()
	var deleteIndex int

	for i, _ := range *s {
		if (*s)[i].id == studentId {
			deleteIndex = i
			break
		}
	}
	fmt.Print("是否继续删除(Y/N):")
	if strings.ToUpper(inputString()) == "Y" {
		result := removeStudentByIndex(deleteIndex, *s)
		return removeStudent(&result)
	}
	return removeStudentByIndex(deleteIndex, *s)
}

func removeStudentByIndex(index int, students []Student) []Student {
	return append(students[:index], students[index+1])
}

func modifyStudent(s *[]Student) {
	fmt.Print("请输入学号")
	studentId := inputInt()
	for i, _ := range *s {
		if studentId == (*s)[i].id {
			fmt.Print("请重新输入学生成绩(语文 数学 英语):")
			(*s)[i].chinese = inputFloat()
			(*s)[i].math = inputFloat()
			(*s)[i].english = inputFloat()

			fmt.Print("是否继续(Y/N):")
			if strings.ToUpper(inputString()) == "Y" {
				modifyStudent(s)
			}
		}
	}
	fmt.Printf("没找到学号为%d的学生信息\n", studentId)
}

func inputTranscript(s *[]Student) []Student {
	for {
	x:
		fmt.Println("请输入学生的学号 姓名 语文成绩 数学成绩 英语成绩")
		student := NewStudent(inputInt(), inputString(), inputFloat(), inputFloat(), inputFloat())
		student.avgScore = avg(*student)
		student.sumScore = sum(*student)
		if !checkRange(*student) {
			fmt.Println("成绩超出范围，请重新输入(成绩范围1-100)")
			goto x
		}
		for _, ss := range *s {
			if ss.id == student.id {
				fmt.Println("学号重复，请重新输入")
				goto x
			}
		}
		*s = append(*s, *student)
		fmt.Print("是否继续输入(输入Y继续):")
		if strings.ToUpper(inputString()) != "Y" {
			break
		}
	}
	return *s
}

func sum(student Student) float64 {
	return student.chinese + student.math + student.english
}

func avg(student Student) float64 {
	return (student.chinese + student.math + student.english) / 3.0
}

func checkRange(student Student) bool {
	if 0 <= student.chinese && student.chinese <= 100 && 0 <= student.math && student.math <= 100 && 0 <= student.english && student.english <= 100 {
		return true
	}
	return false
}

func showTranscript(students []Student) {
	if students == nil {
		fmt.Println("没有学生，请添加学生")
		return
	}
	fmt.Println("学号 姓名 语文 数学 英语 平均分 总成绩")
	for _, student := range students {
		fmt.Println(student.id, " ", student.name, " ", student.chinese, " ", student.math, " ", student.english, " ", student.avgScore, " ", student.sumScore)
	}
}

func filterStudent(s *[]Student) {
	fmt.Println("请输入要查看学生的学号")
	studentId := inputInt()
	for _, student := range *s {
		if student.id == studentId {
			fmt.Println("学号:", student.id, " ",
				"姓名:", student.name, " ",
				"语文:", student.chinese, " ",
				"数学:", student.math, " ",
				"英语:", student.english, " ",
				"平均分:", student.avgScore, " ",
				"总成绩:", student.sumScore)
		}
	}
}

func inputInt() int {
	var inputValue int
	fmt.Scan(&inputValue)
	return inputValue
}

func inputString() string {
	var inputValue string
	fmt.Scan(&inputValue)
	return inputValue
}

func inputFloat() float64 {
	var inputValue float64
	fmt.Scan(&inputValue)
	return inputValue
}

func printMenu() {
	clearScreen()
	fmt.Println("1. 查看学生成绩单")
	fmt.Println("2. 输入学生成绩")
	fmt.Println("3. 修改学生成绩")
	fmt.Println("4. 删除学生成绩")
	fmt.Println("5. 查询学生成绩")
	fmt.Println("6. 对学生总成绩排序")
	fmt.Println("7. 退出系统")
	fmt.Print("请输入你的选择: ")
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
