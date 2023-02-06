package mysql

import (
	"corporeit/backend/model"
	"fmt"
	"go.uber.org/zap"
)

func AddStudent(student *model.Student) error {
	sqlStr := `insert into stu(stuId,class_id,stuName,gender,birthday,user_id) values(?,?,?,?,?,?)`
	_, err := db.Exec(sqlStr, student.StudentId, student.ClassId, student.StudentName, student.Gender, student.Birthday, student.UserId)
	if err != nil {
		fmt.Println(err)
		zap.L().Error("Insert Student err")
	}
	return err
}

func QueryAll() ([]model.Student, error) {
	sqlStr := `select stuId,class_id,stuName,gender,birthday,user_id from stu`
	var students []model.Student
	err := db.Select(&students, sqlStr)
	fmt.Println(students)
	fmt.Println(err)
	if err != nil {
		zap.L().Error("QueryAll err")
		return nil, err
	}
	return students, err
}
