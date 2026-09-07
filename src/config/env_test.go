package config

import (
	"testing"
)

func TestEnv(t *testing.T) {
	InitLoadDotenv()

	settings := NewPortainerSettings()
	if settings == nil {
		t.Fatal("Settings for ApiPortainerSettings is empty!")
	}
}
