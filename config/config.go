package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	InMem       *InMemConfig  `yaml:"in_mem"`
	Worker      *WorkerConfig `yaml:"worker"`
	RoninConfig *RoninConfig  `yaml:"ronin"`
}

type InMemConfig struct {
	Capacity int `yaml:"capacity"`
}

type RoninConfig struct {
	BaseEndPoint string `yaml:"base_end_point"`
}

type WorkerConfig struct {
	Interval int `yaml:"interval"`
}

func Load(filePath string) (*AppConfig, error) {
	if len(filePath) == 0 {
		filePath = os.Getenv("CONFIG_FILE")
	}

	configBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Failed to load config file")
		return nil, err
	}
	configBytes = []byte(os.ExpandEnv(string(configBytes)))

	cfg := &AppConfig{}

	err = yaml.Unmarshal(configBytes, cfg)
	if err != nil {
		log.Printf("Failed to parse config file")
		return nil, err
	}
	log.Printf("config: %+v", cfg)
	log.Printf("======================================")
	log.Printf("in mem config: %+v", cfg.InMem)
	log.Printf("worker config: %+v", cfg.Worker)
	log.Printf("ronin config: %+v", cfg.RoninConfig)

	return cfg, nil
}
