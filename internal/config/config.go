package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost          string  `env:"DB_HOST"`
	DBPort          string  `env:"DB_PORT"`
	DBUser          string  `env:"DB_USER"`
	DBPass          string  `env:"DB_PASS"`
	DBName          string  `env:"DB_NAME"`
	ServerPort      string  `env:"PORT"`
	DatabaseURL     string  `env:"DATABASE_URL"`
	RabbitMQURL     string  `env:"RABBITMQ_URL"`
	MQTTHost        string  `env:"MQTT_HOST"`
	MQTTPort        string  `env:"MQTT_PORT"`
	GeofenceLat     float64 `env:"GEOFENCE_LAT"`
	GeofenceLon     float64 `env:"GEOFENCE_LON"`
	GeofenceRadiusM float64 `env:"GEOFENCE_RADIUS_METERS"`
}

var Cfg Config

func getEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}

func getEnvFloat(key string, defaultVal float64) float64 {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		log.Printf("Warning: Invalid float value for %s, using default value %f. Error: %v", key, defaultVal, err)
		return defaultVal
	}
	return f
}

func Load() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: .env file not loaded, using system env")
	}

	Cfg = Config{
		DBHost:          getEnv("DB_HOST", "localhost"),
		DBPort:          getEnv("DB_PORT", "5432"),
		DBUser:          getEnv("DB_USER", "fleetuser"),
		DBPass:          getEnv("DB_PASS", "fleetpass"),
		DBName:          getEnv("DB_NAME", "fleetdb"),
		ServerPort:      getEnv("PORT", "3000"),
		DatabaseURL:     getEnv("DATABASE_URL", ""),
		RabbitMQURL:     getEnv("RABBITMQ_URL", ""),
		MQTTHost:        getEnv("MQTT_HOST", "tcp://localhost"),
		MQTTPort:        getEnv("MQTT_PORT", "1883"),
		GeofenceLat:     getEnvFloat("GEOFENCE_LAT", -6.2088),
		GeofenceLon:     getEnvFloat("GEOFENCE_LON", 106.8456),
		GeofenceRadiusM: getEnvFloat("GEOFENCE_RADIUS_METERS", 50),
	}
}
