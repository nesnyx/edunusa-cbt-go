package dtos

type LoginStudent struct {
	NIS      string `json:"nis" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginTeacher struct {
	NIK      string `json:"nik" binding:"required"`
	Password string `json:"password" binding:"required"`
}
