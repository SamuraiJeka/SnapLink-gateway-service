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
		LinkGRPC: os.Getenv("LINKGRPC"),
		RedisAddr: os.Getenv("REDISADDR"),
	}
}
