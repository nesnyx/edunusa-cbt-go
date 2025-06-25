package dtos

// Nothing

type ExamTokenUsageRequest struct {
	TokenValueUsed string `json:"token_value_used" binding:"required"`
	ExamID         string `json:"exam_id" binding:"required"`
}
