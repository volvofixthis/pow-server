package config

import (
	"log"
	"time"

	"github.com/caarlos0/env"
)

type AppConfig struct {
	ApiAddress      string        `env:"API_ADDRESS"`
	TCPAddress      string        `env:"TCP_ADDRESS"`
	PowTimeout      time.Duration `env:"POW_TIMEOUT"`
	PowIteration    int           `env:"POW_ITERATION"`
	PowMemory       int           `env:"POW_MEMORY"`
	EmissionSize    int           `env:"EMISSION_SIZE"`
	EmissionDelay   time.Duration `env:"EMISSION_DELAY"`
	ConQueue        int           `env:"CON_QUEUE"`
	ConReadTimeout  time.Duration `env:"CON_READ_TIMEOUT"`
	ConWriteTimeout time.Duration `env:"CON_WRITE_TIMEOUT"`
	WorkerPool      int           `env:"WORKER_POOL"`
}

func NewCfg() *AppConfig {
	cfg := AppConfig{}
	if err := env.Parse(&cfg); err != nil {
		log.Panic(err.Error())
	}
	return &cfg
}
