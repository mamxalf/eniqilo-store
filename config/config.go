package config

import (
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	AppName string `mapstructure:"APP_NAME"`
	AppPort string `mapstructure:"APP_PORT"`
	AppEnv  string `mapstructure:"APP_ENV"`
	AppUrl  string `mapstructure:"APP_URL"`

	CorsAllowCredentials bool     `mapstructure:"CORS_ALLOW_CREDENTIALS"`
	CorsAllowedHeaders   []string `mapstructure:"CORS_ALLOWED_HEADERS"`
	CorsAllowedMethods   []string `mapstructure:"CORS_ALLOWED_METHODS"`
	CorsAllowedOrigins   []string `mapstructure:"CORS_ALLOWED_ORIGINS"`
	CorsEnable           bool     `mapstructure:"CORS_ENABLE"`
	CorsMaxAgeSeconds    int      `mapstructure:"CORS_MAX_AGE_SECONDS"`

	LogLevel   string        `mapstructure:"LOG_LEVEL"`
	JwtSecret  string        `mapstructure:"JWT_SECRET"`
	JwtExpiry  time.Duration `mapstructure:"JWT_EXPIRY"`
	BcryptSalt int           `mapstructure:"BCRYPT_SALT"`

	DbName     string `mapstructure:"DB_NAME"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbUsername string `mapstructure:"DB_USERNAME"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbParams   string `mapstructure:"DB_PARAMS"`
}

var (
	conf Config
	once sync.Once
)

func Get() *Config {
	viper.SetConfigFile(".env")
	viper.ReadInConfig() // Errors here can be handled appropriately.

	// Setting default values
	viper.SetDefault("APP_NAME", "Cat-Social")
	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("APP_ENV", "development")
	viper.SetDefault("APP_URL", "http://localhost:8080")

	viper.SetDefault("CORS_ALLOW_CREDENTIALS", true)
	viper.SetDefault("CORS_ALLOWED_HEADERS", []string{"Accept", "Authorization", "Content-Type"})
	viper.SetDefault("CORS_ALLOWED_METHODS", []string{"GET", "PUT", "POST", "PATCH", "DELETE", "OPTIONS"})
	viper.SetDefault("CORS_ALLOWED_ORIGINS", []string{"*"})
	viper.SetDefault("CORS_ENABLE", true)
	viper.SetDefault("CORS_MAX_AGE_SECONDS", 300)

	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("JWT_SECRET", "abc")
	viper.SetDefault("JWT_EXPIRY", "480m")
	viper.SetDefault("BCRYPT_SALT", 10)

	viper.SetDefault("DB_NAME", "postgres")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_USERNAME", "postgres")
	viper.SetDefault("DB_PASSWORD", "password")
	viper.SetDefault("DB_PARAMS", "sslmode=disable")

	once.Do(func() {
		log.Info().Msg("Service configuration initialized.")
		err := viper.Unmarshal(&conf)
		if err != nil {
			log.Fatal().Err(err)
		}
	})

	return &conf
}
