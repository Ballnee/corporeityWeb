package logic

import (
	"corporeit/backend/dao/mysql"
	"corporeit/backend/model"
)

func ClassGetById(classId uint32) (err error, class []model.Class) {
	err, class = mysql.ClassGetById(classId)
	return
}

func ClassAnalysisById(classId uint32) (error, any) {
	err, class := mysql.ClassAnalysisById(classId)
	return err, class
}
