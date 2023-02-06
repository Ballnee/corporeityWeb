package logic

import (
	"corporeit/backend/dao/mysql"
	"corporeit/backend/model"
	"fmt"
)

func AddStudent(p *model.ParamAddStu, userId uint64) error {
	student := &model.Student{
		StudentId:   p.StudentId,
		ClassId:     p.ClassId,
		StudentName: p.StudentName,
		Gender:      p.Gender,
		Birthday:    p.Birthday,
		UserId:      userId,
	}
	fmt.Println(student)
	return mysql.AddStudent(student)
}

func QueryAll() ([]model.Student, error) {
	return mysql.QueryAll()
}
