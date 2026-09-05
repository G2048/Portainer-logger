package config

import (
	"os"
	"path/filepath"
	"portainer-logger/src/pkg/utils"

	"github.com/joho/godotenv"
)

// Load from .env API_KEY
type ApiPortainerSettings struct {
	ApiKey string
}

// Load from .env API_KEY
func NewApiPortainerSettings() *ApiPortainerSettings {
	return &ApiPortainerSettings{os.Getenv("API_KEY")}
}

// Must be called first!
func InitLoadDotenv() {
	dir := utils.MustResult(os.Getwd())
	// equivalent of ../../..
	for {
		goModPath := filepath.Join(dir, "go.mod")
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
