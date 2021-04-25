package homework

import (
	"log"

	"github.com/pkg/errors"
)

var ErrNoRows = errors.New("no rows")
var ErrInvalidParam = errors.New("invalid param")
var classSet map[string][]Student
var studentSet map[int]Student

type Student struct {
	Id   int
	Name string
}

// 使用map 模拟数据库
func init() {
	u1 := Student{Id: 1, Name: "小明"}
	u2 := Student{Id: 2, Name: "小红"}
	u3 := Student{Id: 3, Name: "小白"}
	u4 := Student{Id: 4, Name: "小绿"}

	classSet = make(map[string][]Student, 0)
	classSet["一年级"] = []Student{u1, u2}
	classSet["二年级"] = []Student{u3, u4}
	classSet["三年级"] = []Student{}

	studentSet = make(map[int]Student, 0)
	studentSet[u1.Id] = u1
	studentSet[u2.Id] = u2
	studentSet[u3.Id] = u3
	studentSet[u4.Id] = u4
}

// 模拟数据库获取数据
// 作为学生集合查询语句
func getStudentsByClass(className string) ([]Student, error) {
	if _, exist := classSet[className]; !exist {
		return nil, ErrNoRows
	}
	return classSet[className], nil
}

// 作为单个学生查询语句
func getStudentById(id int) (Student, error) {
	if id < 0 {
		return Student{}, ErrInvalidParam
	}
	if _, exist := studentSet[id]; !exist {
		return Student{}, ErrNoRows
	}
	return studentSet[id], nil
}

// dao层(模拟获取数据)

// GetStudentsByClass 查询数据集的时候返回不使用wrap包装
func GetStudentsByClass(className string) ([]Student, error) {
	ss, err := getStudentsByClass(className)
	if err != nil && err != ErrNoRows {
		return nil, err
	}

	return ss, nil
}

// GetStudentById 使用wrap包装，附带一些参数信息
func GetStudentById(id int) (Student, error) {
	s, err := getStudentById(id)
	return s, errors.Wrapf(err, "user id: %d", id)
}

// service层(模拟使用数据)

func PrintClassStudentsMsg(className string) {
	ss, err := GetStudentsByClass(className)
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}

	if len(ss) == 0 {
		log.Printf("class: %s no students\n", className)
	}

	for _, v := range ss {
		log.Printf("id: %d,name: %s\n", v.Id, v.Name)
	}
}

func PrintStudentMsg(id int) {
	s, err := GetStudentById(id)
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}

	log.Printf("id: %d,name: %s\n", s.Id, s.Name)
}
