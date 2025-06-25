package repository

import (
	"cbt/internal/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=user_production password=aXRuLWVkdW51c2EtcHNxbC1wcm9kdWN0aW9u dbname=%s port=3499 sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DBHost, cfg.DBName,
	)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Sesuaikan level log GORM
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connection established successfully.")
	sqlDB, err := DB.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(cfg.DBMaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.DBMaxOpenConns)
	sqlDB.SetConnMaxLifetime(cfg.DBConnMaxLifetime)

	// err = migrateTables(DB)
	// if err != nil {
	// 	return nil, fmt.Errorf("database migration failed: %w", err)
	// }
	// log.Println("Database migration completed successfully.")

	return DB, nil
}

// migrateTables menjalankan GORM AutoMigrate untuk semua model
// func migrateTables(db *gorm.DB) error {
// 	err := db.AutoMigrate(
// 		&models.QuestionBank{},       // Tergantung Subject, User (Teacher)
// 		&models.Question{},           // Tergantung QuestionBank, User (Teacher)
// 		&models.Exam{},               // Tergantung Subject, Class, User (Teacher)
// 		&models.ExamQuestion{},       // Tergantung Exam, Question
// 		&models.StudentExamAttempt{}, // Tergantung User (Student), Exam, User (Teacher for grading)
// 		&models.StudentAnswer{},      // Tergantung StudentExamAttempt, ExamQuestion
// 		&models.ExamTokenUsage{},     // Tergantung User (Student), Exam
// 	)
// 	if err != nil {
// 		log.Printf("Error during migration: %v\n", err)
// 		return err
// 	}
// 	return nil
// }
