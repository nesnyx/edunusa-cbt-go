package dtos

type ClassRequest struct {
	ClassName   string `json:"class_name" binding:"required"`
	GradeLevel  string `json:"grade_level" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type ClassResponse struct {
	ID string `json:"id" binding:"required"`
	ClassName   string `json:"class_name" binding:"required"`
	GradeLevel  string `json:"grade_leve" binding:"required"`
	Description string `json:"description" binding:"required"`
}