package logic

import (
	"corporeit/backend/dao/mysql"
	"corporeit/backend/model"
)

func GradeAdd(g *model.ParamAddGrade) error {
	grade := &model.Grade{
		StudentId:        g.StudentId,
		Height:           g.Height,
		Weight:           g.Weight,
		Lungs:            g.Lungs,
		M50:              g.M50,
		SittingForward:   g.SittingForward,
		SitUps:           g.SitUps,
		M50_8:            g.M50_8,
		M800:             g.M800,
		M1000:            g.M1000,
		PullUp:           g.PullUp,
		StandingLongJump: g.StandingLongJump,
		Skips:            g.Skips,
	}
	return mysql.GradeAdd(grade)
}

func QueryAllGrade() (err error, grades []model.Grade) {
	err, grades = mysql.QueryAllGrade()
	return
}
