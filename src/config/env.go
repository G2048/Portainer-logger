package config

import (
	"errors"
	"os"
	"path/filepath"
	"portainer-logger/src/pkg/utils"

	"github.com/joho/godotenv"
)

// Load from .env API_KEY
type PortainerSettings struct {
	ApiKey string
}

// Load from .env API_KEY
func NewPortainerSettings() *PortainerSettings {
	settings := &PortainerSettings{os.Getenv("PORTAINER_API_KEY")}
	if settings.ApiKey == "" {
		panic(errors.New("Cannot find PORTAINER_API_KEY env var!"))
	}
	return settings
}

// Must be called first!
func InitLoadDotenv() {
	dir := utils.MustResult(os.Getwd())
	// equivalent of ../../..
	for {
		goModPath := filepath.Join(dir, ".env")
		// check of exist go.mod
		_, err := os.Stat(goModPath)
		if err == nil {
			break
		}
		dir = filepath.Dir(dir)
	}
	err := godotenv.Load(dir + "/.env")
	if err != nil {
		panic(err)
	}
}
