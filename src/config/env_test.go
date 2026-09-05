package config

import (
	"testing"
)

func TestEnv(t *testing.T) {
	InitLoadDotenv()

	settings := NewApiPortainerSettings()
	if settings == nil {
		t.Fatal("Settings for ApiPortainerSettings is empty!")
	}
}
