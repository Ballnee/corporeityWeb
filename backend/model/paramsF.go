package model

import "time"

// ParamSignUp Signup parameters
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLoginIn Login parameters
type ParamLoginIn struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// param get class
type ParamGetClassById struct {
	ClassId uint32 `json:"classId" binding:"required"`
}

// param get year
type ParamGetYearById struct {
	YearId uint32 `json:"yearId" binding:"required"`
}

// ParamAddStu gender 1 for famale 2 for male
type ParamAddStu struct {
	StudentId   uint32    `json:"studentId" binding:"required"`
	YearId      uint32    `json:"yearId" binding:"required"`
	ClassId     uint32    `json:"classId" binding:"required"`
	StudentName string    `json:"studentName" binding:"required"`
	Gender      uint8     `json:"gender" binding:"required"`
	Birthday    time.Time `json:"birthday" binding:"required"`
}

type ParamAddGrade struct {
	StudentId        uint32  `json:"studentId" binding:"required"`
	Height           uint8   `json:"height" binding:"required"`
	Weight           uint8   `json:"weight" binding:"required"`
	Lungs            uint16  `json:"lungs" binding:"required"`
	M50              float32 `json:"m_50" `
	SittingForward   uint8   `json:"sittingForward"`
	SitUps           uint8   `json:"sitUps" `
	M50_8            float32 `json:"m50_8"`
	M800             float32 `json:"m800"`
	M1000            float32 `json:"m1000"`
	PullUp           uint8   `json:"pullUp"`
	StandingLongJump uint16  `json:"standingLongJump"`
	Skips            uint16  `json:"skips"`
}
