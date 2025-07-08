package middleware

import (
	"cbt/internal/models"
	"cbt/internal/repository"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CheckingTokenExam(db *gorm.DB) gin.HandlerFunc {
	tokenUsageRepo := repository.NewExamTokenUsageRepository(db)
	examRepo := repository.NewExamRepository(db)

	return func(c *gin.Context) {
		currentUser, exists := c.Get(ContextCurrentUser)
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		studentID := currentUser.(ClaimResult).ID

		// 1. Ambil token usage berdasarkan student
		checkingToken, err := tokenUsageRepo.GetByStudent(studentID)

		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			// 2. Jika token belum tercatat (first time masuk), ambil exam_id dari query atau body
			examID := c.Query("examId") // atau dari body/params sesuai implementasi kamu
			if examID == "" {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "examId is required"})
				return
			}

			// 3. Cek exam tersedia
			exam, err := examRepo.GetExamByID(examID)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "exam not found"})
				return
			}

			// 4. Buat token usage baru karena belum ada (anggap mulai ujian)
			newTokenUsage := &models.ExamTokenUsage{
				ExamID:         exam.ID,
				StudentID:      studentID,
				UsageTimestamp: time.Now(),
			}
			_, err = tokenUsageRepo.Create(newTokenUsage)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to create token usage"})
				return
			}

			// lanjutkan request, token valid & baru digunakan
			c.Next()
			return
		} else if err != nil {
			// jika error lainnya
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to check token"})
			return
		}

		// 5. Token sudah ada, cek exam
		exam, err := examRepo.GetExamByID(checkingToken.ExamID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "exam not found"})
			return
		}

		// 6. Cek apakah token sudah expired berdasarkan waktu penggunaan + durasi ujian
		expirationTime := checkingToken.UsageTimestamp.Add(time.Duration(exam.DurationMinutes) * time.Minute)

		if time.Now().After(expirationTime) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
			return
		}

		// 7. Token masih berlaku, lanjutkan
		c.Next()
	}
}
