package config

import (
	// Go Local Packages
	"log"
	"fmt"

	// External Packages
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/rawbytes"
)

func LoadConfig() (*Config, error) {
	k := koanf.New(".")

	if err := k.Load(rawbytes.Provider(DefaultConfig), yaml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	var cfg Config
	// Unmarshal the loaded configuration into the struct
	if err := k.Unmarshal("", &cfg); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &cfg, nil
}