package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment     string // develop, staging, production
	MONGOHost       string
	MONGOPort       string
	MONGOUser       string
	MONGODB         string
	MONGOPassword   string
	POLLServiceHost string
	POLLServicePort string
	LogLevel        string
}

func Load() Config {

	c := Config{}

	c.Environment = cast.ToString(GetOrReturnDefault("ENVIRONMENT", "develop"))
	c.LogLevel = cast.ToString(GetOrReturnDefault("LOG_LEVEL", "debug"))

	c.MONGOHost = cast.ToString(GetOrReturnDefault("MONGO_HOST", "localhost"))
	c.MONGOPort = cast.ToString(GetOrReturnDefault("MONGO_PORT", "5432"))
	c.MONGODB = cast.ToString(GetOrReturnDefault("MONGO_DATABASE", "testdb"))
	c.MONGOPassword = cast.ToString(GetOrReturnDefault("MONGO_PASSWORD", "12321"))
	c.MONGOUser = cast.ToString(GetOrReturnDefault("MONGO_USER", "citizenfour"))

	c.POLLServiceHost = cast.ToString(GetOrReturnDefault("POLL_SERVICE_HOST", "localhost"))
	c.POLLServicePort = cast.ToString(GetOrReturnDefault("POLL_SERVICE_PORT", "1111"))

	return c

}

func GetOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
