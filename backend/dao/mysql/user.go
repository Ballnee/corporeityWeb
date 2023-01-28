package mysql

import (
	"corporeit/backend/model"
	"crypto/md5"
	"encoding/hex"
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
	pwd := encryptPassword(user.Password)
	//Execute the sql statement
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserId, user.Username, pwd)
	if err != nil {
		fmt.Println("InsertUser err")
	}
	return err

}

func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
