package homework

import (
	"fmt"
	"testing"
)

func TestPrintClassStudentsMsg(t *testing.T) {
	// 一年级有学生
	PrintClassStudentsMsg("一年级")
	fmt.Println("----------------")

	// 三年级没有学生
	PrintClassStudentsMsg("三年级")
	fmt.Println("----------------")

	// 不存在一百年级
	PrintClassStudentsMsg("一百年级")
}

func TestPrintStudentMsg(t *testing.T) {
	// 非法学号
	PrintStudentMsg(-1)
	fmt.Println("----------------")

	// 对应学生
	PrintStudentMsg(1)
	fmt.Println("----------------")

	// 不存在该学生
	PrintStudentMsg(100)
}
