package dtos

import (
	"github.com/google/uuid"
)

type ExamRequest struct {
	ExamTitle       string `json:"exam_title" binding:"required"`
	Instructions    string `json:"instructions" binding:"required"`
	SubjectID       string `json:"subject_id" binding:"required"`
	ClassID         string `json:"class_id" binding:"required"`
	StartDatetime   int64  `json:"start_datetime" binding:"required"`
	EndDatetime     int64  `json:"end_datetime" binding:"required"`
	DurationMinutes int    `json:"duration_minutes" binding:"required"`
}

type ExamRequestUpdate struct {
	ExamTitle       string `json:"exam_title" binding:"required"`
	Instructions    string `json:"instructions" binding:"required"`
	ClassID         string `json:"class_id" binding:"required"`
	StartDatetime   int    `json:"start_datetime" binding:"required"`
	EndDatetime     int    `json:"end_datetime" binding:"required"`
	DurationMinutes int    `json:"duration_minutes" binding:"required"`
}

type ExamRequestCreatedByTeacherID struct {
	CreatedByTeacherID uuid.UUID `json:"created_by_teacher_id" binding:"required"`
}

type ExamRequestStudentID struct {
	StudentID string `json:"student_id" binding:"required"`
}

type ExamRequiestByID struct {
	ID uuid.UUID `json:"id" binding:"required"`
}

type ExamResponse struct {
	ID                 uuid.UUID `json:"id"`
	CreatedByTeacherID string    `json:"created_by_teacher_id"`
	AccessTokenExam    string    `json:"access_token_exam"`
	ExamTitle          string    `json:"exam_title"`
	Instruction        string    `json:"instruction"`
	DurationMinutes    int       `json:"duration_minutes"`
	StartDatetime      int64     `json:"start_datetime"`
	EndDatetime        int64     `json:"end_datetime"`
	Status             any       `json:"status"`
}
