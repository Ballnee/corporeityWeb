package logic

import (
	"corporeit/backend/dao/mysql"
	"corporeit/backend/model"
	"corporeit/backend/pkg/jwt"
	"corporeit/backend/pkg/snowflake"
	"errors"

	"go.uber.org/zap"
)

func SignUp(p *model.ParamSignUp) (err error) {
	exists, err := mysql.IsUserExist(p.Username)

	if err != nil {
		zap.L().Error("mysql has err")
		return err
	}
	if exists {
		return errors.New("user already exist")
	}
	//2.create UID
	userId, _ := snowflake.GetID()
	//fmt.Println("snowID", userId)
	//create a user instance
	u := model.User{
		Username: p.Username,
		Password: p.Password,
		UserId:   userId,
	}

	//4.Save into the database
	return mysql.InsertUser(&u)

}

func Login(p *model.ParamLoginIn) (token string, err error) {
	//前端数据和数据库数据对比
	user := &model.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		zap.L().Error("login err")
		return "", err
	}
	return jwt.GenToken(user.UserId, user.Username)

}
