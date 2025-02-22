package configs

type Config struct {
	redisConfig    *RedisConfig
	loggerConfig   *LoggerConfig
	databaseConfig *DatabaseConfig
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	Database int
}

type LoggerConfig struct {
	LogLevel string
}

type DatabaseConfig struct {
	Name         string
	UserName     string
	UserPassword string
	Host         string
	Port         int
}

func NewDatabaseConfig(dbName string, dbUserName string, dbUserPassword string, dbPort int, dbHost string) DatabaseConfig {
	return DatabaseConfig{
		Name:         dbName,
		UserName:     dbUserName,
		UserPassword: dbUserPassword,
		Port:         dbPort,
		Host:         dbHost,
	}
}

func NewRedisConfig(host string, password string, dbCount int, port int) RedisConfig {
	return RedisConfig{
		Host:     host,
		Port:     port,
		Password: password,
		Database: dbCount,
	}
}

func NewLoggerConfig(logLevel string) LoggerConfig {
	return LoggerConfig{LogLevel: logLevel}
}

func New(dbConfig DatabaseConfig, redisConfig RedisConfig, loggerConfig LoggerConfig) Config {
	return Config{
		databaseConfig: &dbConfig,
		redisConfig:    &redisConfig,
		loggerConfig:   &loggerConfig,
	}
}

func (config *Config) RedisConfig() RedisConfig {
	return *config.redisConfig
}

func (config *Config) DatabaseConfig() DatabaseConfig {
	return *config.databaseConfig
}

func (config *Config) LoggerConfig() LoggerConfig {
	return *config.loggerConfig
}

func (redisConfig *RedisConfig) Address() string {
	return redisConfig.Host + ":" + string(redisConfig.Port)
}
