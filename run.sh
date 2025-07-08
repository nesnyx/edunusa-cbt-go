#!/bin/bash

APP_NAME="cbt-go"
GO_FILE="./cmd/api/main.go"

echo "🚀 Build aplikasi Go..."
go build -o $APP_NAME $GO_FILE || { echo "❌ Gagal build"; exit 1; }

echo "✅ Build berhasil. Menjalankan dengan PM2..."
pm2 restart "$APP_NAME"
pm2 start ./$APP_NAME --name "$APP_NAME"
