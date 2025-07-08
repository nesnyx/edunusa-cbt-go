package dtos

type StudentAnswerRequest struct {
	AnswerData           string `json:"answer_data" binding:"required"`
	StudentExamAttemptId string `json:"student_exam_attempt_id" binding:"required"`
	ExamQuestionID       string `json:"exam_question_id" binding:"required"`
}

type GetStudentAnswerByQuestionAndStudentAttempt struct {
	StudentID    string `json:"student_id" binding:"required"`
	ExamQuestion string `json:"exam_question" binding:"required"`
}
