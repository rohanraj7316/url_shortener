package services

import (
	"log"
)

// InitService initialize the service
func InitService() {
	defer func() {
		if err := recover(); err != nil {
			// TODO: need to add logging.
			log.Printf("[Error] service initialization: %v", err)
		}
	}()

	err := LoadConfig()
	if err != nil {
		panic(err)
	}

	err = StartServices()
	if err != nil {
		panic(err)
	}

	err = HealthCheck()
	if err != nil {
		panic(err)
	}
}

// LoadConfig loading configs
func LoadConfig() error {
	// TODO: load .env files
	return nil
}

// StartServices start all the services
func StartServices() error {

	var err error

	// TODO: load the config and pass along all the variable

	err = MongoConnection()
	if err != nil {
		return err
	}

	RedisConnection()

	return nil
}

// HealthCheck health check for all services
func HealthCheck() error {
	var err error

	err = MongoConnectionHealthCheck()
	if err != nil {
		return err
	}

	err = RedisConnectionHealthCheck()
	if err != nil {
		return err
	}

	return nil
}
