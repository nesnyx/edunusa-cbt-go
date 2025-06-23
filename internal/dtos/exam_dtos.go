package dtos

import (
	"time"

	"github.com/google/uuid"
)

type ExamRequest struct {
	ExamTitle          string    `json:"exam_title" binding:"required"`
	Instructions       string    `json:"instructions" binding:"required"`
	SubjectID          uuid.UUID `json:"subject_id" binding:"required"`
	ClassID            uuid.UUID `json:"class_id" binding:"required"`
	CreatedByTeacherID uuid.UUID `json:"created_by_teacher_id" binding:"required"`
	StartDatetime      time.Time `json:"start_datetime" binding:"required"`
	EndDatetime        time.Time `json:"end_datetime" binding:"required"`
	DurationMinutes    int       `json:"duration_minutes" binding:"required"`
}

type ExamRequestUpdate struct {
	ExamTitle       string    `json:"exam_title" binding:"required"`
	Instructions    string    `json:"instructions" binding:"required"`
	ClassID         string    `json:"class_id" binding:"required"`
	StartDatetime   time.Time `json:"start_datetime" binding:"required"`
	EndDatetime     time.Time `json:"end_datetime" binding:"required"`
	DurationMinutes int       `json:"duration_minutes" binding:"required"`
}

type ExamRequestCreatedByTeacherID struct {
	CreatedByTeacherID uuid.UUID `json:"created_by_teacher_id" binding:"required"`
}

type ExamRequiestByID struct {
	ID uuid.UUID `json:"id" binding:"required"`
}

type ExamResponse struct {
	ID                 uuid.UUID `json:"id"`
	CreatedByTeacherID uuid.UUID `json:"created_by_teacher_id"`
	AccessTokenExam    string    `json:"access_token_exam"`
	ExamTitle          string    `json:"exam_title"`
}
