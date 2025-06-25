package repositoryextention

import (
	"cbt/extentions/configExtention"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDBExtention(cfg *configExtention.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=user_production password=aXRuLWVkdW51c2EtcHNxbC1wcm9kdWN0aW9u dbname=%s port=3499 sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DBHost, cfg.DBName,
	)
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	cfg.DBUser,
	// 	cfg.DBPassword,
	// 	cfg.DBHost,
	// 	cfg.DBPort,
	// 	cfg.DBName,
	// )

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

	// Menjalankan AutoMigrate

	// err = migrateTables(DB)
	// if err != nil {
	// 	return nil, fmt.Errorf("database migration failed: %w", err)
	// }
	// log.Println("Database migration completed successfully.")

	return DB, nil
}

// func migrateTables(db *gorm.DB) error {
// 	err := db.AutoMigrate(
// 		&models.Class{},
// 		&models.Role{},
// 		&models.Student{},
// 		&models.Teacher{},
// 		&models.Subject{},
// 		&models.HasRole{},
// 	)
// 	if err != nil {
// 		log.Printf("Error during migration: %v\n", err)
// 		return err
// 	}
// 	return nil
// }
