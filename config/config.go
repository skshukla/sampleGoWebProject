package config

import rateLimitConfig "github.com/skshukla/sampleRateLimit/config"

type AppConfig struct {
	Database struct{
		Host string
		Port string
		Username string
		Password string
		DBName string
	}
	Server struct {
		Host string
		Port string
	}
	RateLimitConfig rateLimitConfig.RateLimitConfig `yaml:"rateLimitConfig"`
}
