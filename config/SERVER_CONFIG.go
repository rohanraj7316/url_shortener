package config

type serverConfigSchema struct {
	Port       uint
	BodyLimit  string
	CorsOrigin string
	CorsMethod string
}

func validate() (bool, error) {
	return false, nil
}

// GetEnvValue return env value of the given key
func GetEnvValue(key string) string {
	return ""
}
