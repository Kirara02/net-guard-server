package config

import (
	"fmt"
	"os"
	"strconv"

	"NetGuardServer/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     int
}

// JWTConfig holds JWT configuration
type JWTConfig struct {
	Secret             string
	AccessTokenExpDays int
}

type FirebaseConfig struct {
	FirebaseServiceAccountPath string
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port string
}

// Config holds all application configurations
type Config struct {
	Database                   DatabaseConfig
	JWT                        JWTConfig
	Firebase                   FirebaseConfig
	FirebaseServiceAccountPath string
	Server                     ServerConfig
	DB                         *gorm.DB
}

// AppConfig is the global configuration instance
var AppConfig Config

// LoadConfig loads configuration from environment variables and Firebase JSON
func LoadConfig() error {
	// Load database configuration from environment variables
	AppConfig.Database.Host = getEnv("DB_HOST", "localhost")
	AppConfig.Database.User = getEnv("DB_USER", "kirara")
	AppConfig.Database.Password = getEnv("DB_PASSWORD", "")
	AppConfig.Database.Name = getEnv("DB_NAME", "netguard_db")
	portStr := getEnv("DB_PORT", "5432")
	AppConfig.Database.Port, _ = strconv.Atoi(portStr)

	// Load JWT configuration from environment variables
	AppConfig.JWT.Secret = getEnv("JWT_SECRET", "kiraraberntein_adkjdkfhdjfhudhd")
	expStr := getEnv("ACCESS_TOKEN_EXP_DAYS", "7")
	AppConfig.JWT.AccessTokenExpDays, _ = strconv.Atoi(expStr)

	AppConfig.Firebase.FirebaseServiceAccountPath = getEnv("FIREBASE_SERVICE_ACCOUNT_PATH", "config/netguard-7b734-9c58282275ac.json")

	// Load server configuration from environment variables
	AppConfig.Server.Port = getEnv("PORT", "8080")

	// Load Firebase service account path
	AppConfig.FirebaseServiceAccountPath = getEnv("FIREBASE_SERVICE_ACCOUNT_PATH", "config/netguard-7b734-9c58282275ac.json")

	// Initialize database connection
	if err := InitDatabase(); err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}

	return nil
}

// InitDatabase initializes the database connection and runs auto-migration
func InitDatabase() error {
	// Debug: print configuration
	fmt.Printf("Connecting to database: host=%s user=%s password='%s' dbname=%s port=%d\n",
		AppConfig.Database.Host,
		AppConfig.Database.User,
		AppConfig.Database.Password,
		AppConfig.Database.Name,
		AppConfig.Database.Port,
	)

	// Use URL format for PostgreSQL DSN
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		AppConfig.Database.User,
		AppConfig.Database.Password,
		AppConfig.Database.Host,
		AppConfig.Database.Port,
		AppConfig.Database.Name,
	)

	fmt.Printf("DSN: %s\n", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	AppConfig.DB = db

	// Test the connection
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	fmt.Println("✅ Database connection successful")

	// Auto-migrate the schema
	if err := db.AutoMigrate(
		&models.User{},
		&models.Server{},
		&models.ServerDownHistory{},
	); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	fmt.Println("✅ Database migration completed")
	return nil
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
