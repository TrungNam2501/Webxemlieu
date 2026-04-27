package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// Config gom toàn bộ cấu hình runtime của backend.
type Config struct {
	Port               string
	GinMode            string
	FrontendDist       string
	DBHost             string
	DBPort             string
	DBUser             string
	DBPassword         string
	DBName             string
	DBSSLMode          string
	CORSAllowedOrigins []string
}

// Load đọc biến môi trường (kèm file .env nếu có) và trả về Config.
func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("config: không tìm thấy .env (%v) – dùng biến môi trường hệ thống", err)
	}

	cfg := &Config{
		Port:         getEnv("PORT", "8080"),
		GinMode:      getEnv("GIN_MODE", "debug"),
		FrontendDist: getEnv("FRONTEND_DIST", "../frontend/dist"),
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       getEnv("DB_PORT", "5432"),
		DBUser:       getEnv("DB_USER", "postgres"),
		DBPassword:   getEnv("DB_PASSWORD", ""),
		DBName:       getEnv("DB_NAME", "postgres"),
		DBSSLMode:    getEnv("DB_SSLMODE", "disable"),
	}

	origins := getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:5173")
	for _, o := range strings.Split(origins, ",") {
		o = strings.TrimSpace(o)
		if o != "" {
			cfg.CORSAllowedOrigins = append(cfg.CORSAllowedOrigins, o)
		}
	}

	return cfg
}

// PostgresDSN trả về connection string tương thích với gorm/postgres driver.
func (c *Config) PostgresDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode,
	)
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}
