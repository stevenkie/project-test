package util

import "os"

func GetEnv() string {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}
	return env
}
