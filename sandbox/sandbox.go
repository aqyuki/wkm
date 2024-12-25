package sandbox

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var ErrAlreadyExits = errors.New("sandbox directory already exists")

func Create(rootDir, title string) (string, error) {
	sandboxDir := filepath.Join(rootDir, title)

	if _, err := os.Stat(sandboxDir); err == nil {
		return "", ErrAlreadyExits
	} else if !os.IsNotExist(err) {
		return "", fmt.Errorf("failed to check sandbox directory because of %w", err)
	}

	if err := os.MkdirAll(sandboxDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create sandbox directory because of %w", err)
	}
	return sandboxDir, nil
}

func List(rootDir string) ([]string, error) {
	entries, err := os.ReadDir(rootDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read sandbox directory because of %w", err)
	}

	var directories []string
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		directories = append(directories, entry.Name())
	}
	return directories, nil
}

func Remove(rootDir, title string) error {
	sandboxDir := filepath.Join(rootDir, title)
	if err := os.RemoveAll(sandboxDir); err != nil {
		return fmt.Errorf("failed to remove sandbox directory because of %w", err)
	}
	return nil
}
