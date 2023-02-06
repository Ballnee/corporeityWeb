package model

import "time"

type User struct {
	UserId   uint64 `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type Student struct {
	StudentId   uint32    `db:"stuId"`
	ClassId     uint32    `db:"class_id"`
	StudentName string    `db:"stuName"`
	Gender      uint8     `db:"gender"`
	Birthday    time.Time `db:"birthday"`
	UserId      uint64    `db:"user_id"`
}
