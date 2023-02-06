package model

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

// ParamAddStu gender 1 for famale 2 for male
type ParamAddStu struct {
	StudentId   uint32 `json:"studentId" binding:"required"`
	ClassId     uint32 `json:"classId" binding:"required"`
	StudentName string `json:"studentName" binding:"required"`
	Gender      uint8  `json:"gender" binding:"required"`
	Birthday    int64  `json:"birthday" binding:"required"`
}
