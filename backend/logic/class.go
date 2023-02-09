package logic

import (
	"corporeit/backend/dao/mysql"
	"corporeit/backend/model"
)

func ClassGetById(classId uint16) (err error, class []model.Class) {
	err, class = mysql.ClassGetById(classId)
	return
}
