package config

import (
	"errors"
	"os"
	"path/filepath"
)

var (
	ErrEmptyRoot = errors.New("$WKM_ROOT is now allowed to be empty")
)

type Config struct {
	// Root is a path to the root directory of the sandbox
	Root string
}

func (c Config) MustAbsRoot() string {
	abs, err := filepath.Abs(c.Root)
	if err != nil {
		panic(err)
	}
	return abs
}

func (c Config) Validate() error {
	if c.Root == "" {
		return ErrEmptyRoot
	}

	return nil
}

func LoadConfig() (*Config, error) {
	config := &Config{
		Root: os.Getenv("WKM_ROOT"),
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}
	return config, nil
}
