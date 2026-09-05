package portainer

import (
	"portainer-logger/src/config"
	"testing"
)

func TestPortainerApi(t *testing.T) {
	config.InitLoadDotenv()
	settings := config.NewApiPortainerSettings()
	t.Logf("%#+v\n", settings)
	if settings == nil {
		t.Fail()
	}
	if settings.ApiKey == "" {
		t.Fail()
	}

	client := NewPortainerApi(settings.ApiKey)
	res, err := client.Containers(1)
	t.Log(res.Status)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if res.Status != 200 {
		t.Fail()
	}
	t.Logf("%s", res.DecodeBody())
}
