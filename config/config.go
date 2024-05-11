package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"strconv"
	"sync"
	"time"
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
	err := viper.ReadInConfig()

	if err != nil {
		log.Error().Err(err).Msg("Failed reading config file")
		appName := os.Getenv("APP_NAME")
		if appName == "" {
			appName = "Eniqilo Store"
		}
		appPort := os.Getenv("APP_PORT")
		if appPort == "" {
			appPort = "8080"
		}
		logLevel := os.Getenv("LOG_LEVEL")
		if logLevel == "" {
			logLevel = "info"
		}
		jwtExpiry := os.Getenv("JWT_EXPIRY")
		if jwtExpiry == "" {
			jwtExpiry = "480m"
		}
		jwtExpired, _ := time.ParseDuration(jwtExpiry)
		corsEnable := os.Getenv("CORS_ENABLE")
		if corsEnable == "" {
			corsEnable = "true"
		}
		corsEnabler, _ := strconv.ParseBool(corsEnable)
		corsMaxAgeSeconds := os.Getenv("CORS_MAX_AGE_SECONDS")
		if corsMaxAgeSeconds == "" {
			corsMaxAgeSeconds = "600"
		}
		corsMaxAge, _ := strconv.Atoi(corsMaxAgeSeconds)
		// must on env
		bcryptSalt, err := strconv.Atoi(os.Getenv("BCRYPT_SALT"))
		if err != nil {
			log.Fatal().Err(err).Msg("Failed parsing BCRYPT_SALT")
		}
		config := &Config{
			AppName:              appName,
			AppPort:              appPort,
			CorsAllowCredentials: true,
			CorsAllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
			CorsAllowedMethods:   []string{"GET", "PUT", "POST", "PATCH", "DELETE", "OPTIONS"},
			CorsAllowedOrigins:   []string{"*"},
			LogLevel:             logLevel,
			CorsEnable:           corsEnabler,
			CorsMaxAgeSeconds:    corsMaxAge,
			JwtExpiry:            jwtExpired,
			// must on env - no default value
			BcryptSalt: bcryptSalt,
			JwtSecret:  os.Getenv("JWT_SECRET"),
			DbName:     os.Getenv("DB_NAME"),
			DbPort:     os.Getenv("DB_PORT"),
			DbHost:     os.Getenv("DB_HOST"),
			DbUsername: os.Getenv("DB_USERNAME"),
			DbPassword: os.Getenv("DB_PASSWORD"),
			DbParams:   os.Getenv("DB_PARAMS"),
		}
		return config
	} else {
		once.Do(func() {
			log.Info().Msg("Service configuration initialized.")
			err := viper.Unmarshal(&conf)
			if err != nil {
				log.Fatal().Err(err)
			}
		})
		return &conf
	}
}
