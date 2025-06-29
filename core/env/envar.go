package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

var once sync.Once

// loadEnv ensures .env is loaded only once
func loadEnv() {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Println("Warning: Could not load .env file, relying on system environment variables")
		}
	})
}

// GetEnvString returns the value of the environment variable as a string.
// If the variable is not present, it returns an empty string.
func GetEnvString(key string) string {
	loadEnv()
	return os.Getenv(key)
}

// GetEnvInt returns the value of the environment variable as an integer.
// If the variable is not present or cannot be converted to an integer, it returns 0.
func GetEnvInt(key string) int {
	loadEnv()
	val := os.Getenv(key)
	if val == "" {
		return 0
	}
	intVal, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return intVal
}

// GetEnvBool returns the value of the environment variable as a boolean.
// If the variable is not present or cannot be converted to a boolean, it returns false.
func GetEnvBool(key string) bool {
	loadEnv()
	val := os.Getenv(key)
	if val == "" {
		return false
	}
	boolVal, err := strconv.ParseBool(val)
	if err != nil {
		return false
	}
	return boolVal
}

// GetEnvMap returns the value of the environment variable as a map (array) of strings.
// The environment variable should be a comma-separated list of key-value pairs, e.g., "key1=value1,key2=value2".
// If the variable is not present, it returns an empty map.
func GetEnvMap(key string) map[string]string {
	loadEnv()
	val := os.Getenv(key)
	if val == "" {
		return make(map[string]string)
	}

	pairs := strings.Split(val, ",")
	result := make(map[string]string)

	for _, pair := range pairs {
		// Split each pair into key and value
		kv := strings.Split(pair, "=")
		// If the pair is not in the format "key=value", return an empty map
		if len(kv) != 2 {
			return make(map[string]string)
		}
		result[kv[0]] = kv[1]
	}

	return result
}

// GetEnvArray retrieves an environment variable as a []string, splitting by commas,
// and removes all empty strings from the result.
func GetEnvArray(key string) []string {
	loadEnv()
	val := os.Getenv(key)
	if val == "" {
		return []string{}
	}

	parts := strings.Split(val, ",")

	// Filter out empty strings
	var result []string
	for _, part := range parts {
		if part != "" {
			result = append(result, part)
		}
	}

	return result
}
