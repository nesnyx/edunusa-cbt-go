package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error

	// Cek environment variable GIN_MODE untuk menentukan mode logger
	// Defaultnya "development" jika tidak di-set
	env := os.Getenv("GIN_MODE")

	var config zap.Config
	if env == "production" {
		config = zap.NewProductionConfig()
		// Di produksi, kita mungkin ingin level log INFO ke atas
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	} else {
		config = zap.NewDevelopmentConfig()
		// Di development, gunakan format console yang berwarna
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		// Tampilkan semua level log di development
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	// Build logger
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

// Sediakan fungsi helper untuk mengakses logger global
func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	log.Warn(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	log.Fatal(message, fields...)
}
