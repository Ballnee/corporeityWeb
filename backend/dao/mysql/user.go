package mysql

import (
	"corporeit/backend/model"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
)

func IsUserExist(uname string) (bool, error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, uname); err != nil {

		return false, err
	}
	return count > 0, nil
}

var secret = "saxafaaxa"

func InsertUser(user *model.User) (err error) {
	//Password encryption
	pwd := EncryptPassword(user.Password)
	//Execute the sql statement
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserId, user.Username, pwd)
	if err != nil {
		fmt.Println("InsertUser err")
	}
	return err

}

func Login(user *model.User) (err error) {
	enPwd := EncryptPassword(user.Password)
	sqlStr := `select user_id,username,password from user where username=?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return errors.New("username not exists")
	}

	if err != nil {
		return errors.New("databases err")
	}

	if enPwd != user.Password {
		return errors.New("password err")
	}
	return nil
}

func EncryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
