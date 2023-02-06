package mysql

import (
	"corporeit/backend/model"
	"fmt"
	"go.uber.org/zap"
	"time"
)

func AddStudent(student *model.Student) error {
	sqlStr := `insert into stu(stuId,class_id,stuName,gender,birthday,user_id) values(?,?,?,?,?,?)`
	_, err := db.Exec(sqlStr, student.StudentId, student.ClassId, student.StudentName, student.Gender, time.Unix(student.Birthday, 0), student.UserId)
	if err != nil {
		fmt.Println(err)
		zap.L().Error("Insert Student err")
	}
	return err
}
