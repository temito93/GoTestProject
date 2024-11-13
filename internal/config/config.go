package config

import "os"

type Config struct {
	URI  string
	Name string
}

func MongoConfig() *Config {
	return &Config{
		URI:  os.Getenv("MONGODB_URI"),
		Name: os.Getenv("MONGO_DB"),
	}
}
