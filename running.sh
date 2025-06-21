#!/bin/bash

APP_NAME="cbt-go"
GO_FILE="./cmd/api/main.go"
BINARY="./$APP_NAME"

echo "🚀 Build aplikasi Go..."
go build -o $APP_NAME $GO_FILE

if [ $? -ne 0 ]; then
    echo "❌ Gagal membangun aplikasi."
    exit 1
fi

echo "✅ Build berhasil: $BINARY"
./$APP_NAME

