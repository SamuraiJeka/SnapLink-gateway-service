package config

import "os"


type Config struct {
	AuthGRPC string
	LinkGRPC string
	RedisAddr string
}


func Load() *Config {
	return &Config{
		AuthGRPC: os.Getenv("AUTH_GRPC"),
		LinkGRPC: os.Getenv("LINKGRPC"),
		RedisAddr: os.Getenv("REDISADDR"),
	}
}
