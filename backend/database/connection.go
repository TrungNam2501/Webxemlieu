package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB là handle GORM dùng chung cho toàn ứng dụng.
var DB *gorm.DB

// Connect mở kết nối tới PostgreSQL bằng DSN cho trước và gán vào biến DB.
func Connect(dsn string) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("không kết nối được DB: %w", err)
	}
	DB = db
	return nil
}
