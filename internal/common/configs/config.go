package configs

import "sync"

var (
	instance *Config
	once     sync.Once
)

type Config struct {
}

func Get() *Config {
	return instance
}

func Init() *Config {
	once.Do(func() {
		instance = &Config{}
	})
	return instance
}
