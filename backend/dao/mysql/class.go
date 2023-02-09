package mysql

import (
	"corporeit/backend/model"
	"fmt"
)

//func ClassGetById(classId uint16) (err error, class []model.Class) {
//	sqlStrGetStuId := `select stuId,stuName,gender,birthday from stu where class_id = ?`
//	var stuIds []uint32
//	var stus []model.Student
//	err = db.Select(&stus, sqlStrGetStuId, classId)
//	if err != nil {
//		zap.L().Error("ClassGetById sqlStr 1 err")
//		fmt.Println(err)
//		return err, nil
//	}
//	for _, v := range stus {
//		stuIds = append(stuIds, v.StudentId)
//	}
//	//fmt.Println("stuID--", stuIds)
//	var grades []model.Grade
//	sqlStr := `select stuId,height,weight,lungs,50M,sittingForward,50M8,sitUpsPerMin,skipsPerMin,standingLongJump,800M,1000M,pullUp from grades where  stuId IN (?)`
//	query, args, err := sqlx.In(sqlStr, stuIds)
//	if err != nil {
//		zap.L().Error("ClassGetById sqlStr 2 err")
//		fmt.Println(err)
//		return err, nil
//	}
//	query = db.Rebind(query)
//	err = db.Select(&grades, query, args...)
//
//	//fmt.Println(stus)
//	//fmt.Println("grades", grades)
//	for i := 0; i < len(stus); i++ {
//		tempStu := model.Class{}
//		tempStu.Grade.StudentId = stus[i].StudentId
//		tempStu.StudentName = stus[i].StudentName
//		tempStu.Gender = stus[i].Gender
//		tempStu.Birthday = stus[i].Birthday
//		for j := 0; j < len(grades); j++ {
//			if grades[j].StudentId == stus[i].StudentId {
//				tempStu.Grade = grades[j]
//			}
//		}
//		class = append(class, tempStu)
//	}
//
//	//fmt.Println("class---", class)
//	return
//}

func ClassGetById(classId uint16) (err error, class []model.Class) {
	sqlStr := `select t.stuId,stuName,gender,birthday,COALESCE(height,0) as height,ifnull(weight,0) as weight,ifnull(lungs,0) as lungs,ifnull(50M,0) as 50M,ifnull(sittingForward,0) as sittingForward,ifnull(50M8,0) as 50M8,ifnull(sitUpsPerMin,0) as sitUpsPerMin,ifnull(skipsPerMin,0) as skipsPerMin,ifnull(standingLongJump,0) as standingLongJump,ifnull(800M,0) as 800M,ifnull(1000M,0) as 1000M,ifnull(pullUp,0) as pullUp
from stu t left join corporeity.grades g on t.stuId = g.stuId where t.class_id = ?;`
	err = db.Select(&class, sqlStr, classId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(class)
	return
}
