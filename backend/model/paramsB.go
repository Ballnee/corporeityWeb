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
	YearId      uint32    `db:"year_id"`
	StudentName string    `db:"stuName"`
	Gender      uint8     `db:"gender"`
	Birthday    time.Time `db:"birthday"`
	UserId      uint64    `db:"user_id"` //相当于体育老师的id
}

type Grade struct {
	StudentId        uint32  `db:"stuId"`
	Height           uint8   `db:"height"`
	Weight           uint8   `db:"weight"`
	Lungs            uint16  `db:"lungs"`
	M50              float32 `db:"50M"`
	SittingForward   uint8   `db:"sittingForward"`
	SitUps           uint8   `db:"sitUpsPerMin"`
	M50_8            float32 `db:"50M8"`
	M800             float32 `db:"800M"`
	M1000            float32 `db:"1000M"`
	PullUp           uint8   `db:"pullUp"`
	StandingLongJump uint16  `db:"standingLongJump"`
	Skips            uint16  `db:"skipsPerMin"`
}

type Class struct {
	StudentName      string    `db:"stuName"`
	Gender           uint8     `db:"gender"`
	Birthday         time.Time `db:"birthday"`
	StudentId        uint32    `db:"stuId"`
	Height           uint8     `db:"height"`
	Weight           uint8     `db:"weight"`
	Lungs            uint16    `db:"lungs"`
	M50              float32   `db:"50M"`
	SittingForward   uint8     `db:"sittingForward"`
	SitUps           uint8     `db:"sitUpsPerMin"`
	M50_8            float32   `db:"50M8"`
	M800             float32   `db:"800M"`
	M1000            float32   `db:"1000M"`
	PullUp           uint8     `db:"pullUp"`
	StandingLongJump uint16    `db:"standingLongJump"`
	Skips            uint16    `db:"skipsPerMin"`
}

// class
type Ranking struct {
	NonConformingNums int `db:"Non_conformingNums"`
	QualifiedNums     int `db:"qualifiedNums"`
	GoodNums          int `db:"goodNums"`
	Outstanding       int `db:"outstandingNums"`
}
