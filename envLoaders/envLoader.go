package envLoaders

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/nimishgj/go-config/configs"
	"go.uber.org/zap"
)

type EnvLoader struct {
	envFileName string
	logger      *zap.Logger
}

func New(envFileName string, logger *zap.Logger) *EnvLoader {
	return &EnvLoader{
		envFileName: envFileName,
		logger:      logger,
	}
}

func (e *EnvLoader) LoadEnvToConfig() configs.Config {
	err := godotenv.Load(e.envFileName)
	if err != nil {
		e.logger.Error("Error loading env file",
			zap.String("envFileName", e.envFileName),
			zap.Error(err))
		return configs.Config{}
	}

	redisConfig := configs.NewRedisConfig(
		getEnv("REDIS_HOST", "localhost:6379"),
		getEnv("REDIS_PASSWORD", ""),
		0,
		getEnvAsInt("REDIS_PORT", 6379))

	loggerConfig := configs.NewLoggerConfig(
		getEnv("LOG_LEVEL", "info"))

	databaseConfig := configs.NewDatabaseConfig(
		getEnv("DB_NAME", "test_db"),
		getEnv("DB_USER_NAME", "admin"),
		getEnv("DB_USER_PASSWORD", ""),
		getEnvAsInt("DB_PORT", 5432),
		getEnv("DB_HOST", "localhost"))

	return configs.New(databaseConfig, redisConfig, loggerConfig)
}

func getEnv(key, fallback string) string {
	if envValue, ok := os.LookupEnv(key); ok {
		return envValue
	}
	return fallback
}

func getEnvAsInt64(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		envValueAsInt64, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			return envValueAsInt64
		}
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		envValueAsInt, err := strconv.Atoi(value)
		if err == nil {
			return envValueAsInt
		}
	}
	return fallback
}
