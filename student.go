package main

import fmt1 "fmt"

// 描述学生信息的结构体
type student struct {
	id    int
	name  string
	class string
}

// newStudent 是student类型的构造函数
func newStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

// 学员管理的类型
type studentMgr struct {
	allStudents []*student
}

// newStudentMgr 是studentMgr的构造函数
func newStudentMgr() *studentMgr {
	return &studentMgr{
		allStudents: make([]*student, 0, 100),
	}
}

// 1. 添加学生的方法
func (s *studentMgr) addStudent(newStu *student) {
	s.allStudents = append(s.allStudents, newStu)
}

// 2. 编辑学生的方法
func (s *studentMgr) modifyStudent(newStu *student) {

	for i, v := range s.allStudents {
		if newStu.id == v.id {
			s.allStudents[i] = newStu
			return
		}
	}
	// 如果走到这里，说明系统中不存在这个学生的信息
	fmt1.Printf("输入的学生信息有误，系统中不存在%d的信息\n", newStu.id)
}

// 3. 展示所有学生的方法
func (s *studentMgr) showStudent() {
	//fmt1.Printf("%#v\n",s.allStudents)
	for _, v := range s.allStudents {
		fmt1.Printf("学号: %d 姓名: %s 班级: %s\n", v.id, v.name, v.class)
	}
}
