package config

import "os"


type Config struct {
	HttpAddr string

	AuthGRPC string
	LinkGRPC string
	RedisAddr string
}


func Load() *Config {
	return &Config{
		HttpAddr: os.Getenv("HTTP_ADDR"),
		AuthGRPC: os.Getenv("AUTH_GRPC"),
		LinkGRPC: os.Getenv("LINK_GRPC"),
		RedisAddr: os.Getenv("REDIS_ADDR"),
	}
}
