package env

import "os"

func GetEnvOrDefault(env, def string) string {
	if v := os.Getenv(env); v != "" {
		return v
	}
	return def
}

func IsCloud() bool {
	_, ok := os.LookupEnv("CQ_CLOUD")
	return ok
}
