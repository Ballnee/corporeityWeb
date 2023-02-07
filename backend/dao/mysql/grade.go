package mysql

import (
	"corporeit/backend/model"
	"fmt"
	"go.uber.org/zap"
)

func GradeAdd(grade *model.Grade) error {
	sqlStr := `insert into grades(stuId,height,weight,lungs,50M,sittingForward,50M8,sitUpsPerMin,skipsPerMin
				,standingLongJump,800M,1000M,pullUp) values(?,?,?,?,?,?,?,?,?,?,?,?,?)`
	_, err := db.Exec(sqlStr, grade.StudentId, grade.Height, grade.Weight, grade.Lungs, grade.M50, grade.SittingForward, grade.M50_8, grade.SitUps, grade.Skips, grade.StandingLongJump, grade.M800, grade.M1000, grade.PullUp)
	if err != nil {
		zap.L().Error("mysql GradeAdd err")
		fmt.Println(err)
		return err
	}
	return err
}

func QueryAllGrade() (err error, grades []model.Grade) {
	sqlStr := `select stuId,height,weight,lungs,50M,sittingForward,50M8,sitUpsPerMin,skipsPerMin,standingLongJump,800M,1000M,pullUp from grades`
	err = db.Select(&grades, sqlStr)
	if err != nil {
		fmt.Println(grades)
		fmt.Println(err)
		zap.L().Error("QueryAllGrade databases err")
		return
	}
	return
}
