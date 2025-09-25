package config

import (
	"fmt"
	"time"

	"github.com/ak-repo/ecommerce-gin/internal/common/utils"
	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		Port string
		Host string
	}

	Database struct {
		Host     string
		Port     string
		DBName   string
		User     string
		Password string
		SSLMode  string
	}

	JWT struct {
		SecretKey         string
		AccessExpiration  time.Duration
		RefreshExpiration time.Duration
	}
}

// Load env varaibles
func Load() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	config := &Config{}

	config.Server.Host = utils.GetEnv("SERVER_HOST", "8080")
	config.Server.Port = utils.GetEnv("SERVER_PORT", "0.0.0.0")

	config.Database.Host = utils.GetEnv("DB_HOST", "localhost")
	config.Database.Port = utils.GetEnv("DB_PORT", "5432")
	config.Database.User = utils.GetEnv("DB_USER", "ak")
	config.Database.Password = utils.GetEnv("DB_PASSWORD", "4455")
	config.Database.DBName = utils.GetEnv("DB_NAME", "users_db")
	config.Database.SSLMode = utils.GetEnv("DB_SSLMODE", "disable")

	config.JWT.SecretKey = utils.GetEnv("JWT_SECRET", "your-secret-key")
	config.JWT.AccessExpiration = time.Minute * 10 // 5 minutes
	config.JWT.RefreshExpiration = time.Hour * 168 // 7 days

	return config, nil
}

// DB dsn
func (c *Config) GetDSN() string {

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.DBName,
		c.Database.SSLMode)
}

func (c *Config) ServerAddress() string {
	return c.Server.Host + ":" + c.Server.Port
}
