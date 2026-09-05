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

	// strBody, _ := res.DecodeBodyString()
	// fmt.Println(strBody)

	var decodedBody []ContainersResponse
	_, erro := res.DecodeBodyStruct(&decodedBody)
	if erro != nil {
		t.Error(erro)
	}

	// t.Logf("%#v", decodedBody)

	for _, row := range decodedBody {
		if row.State == Running {
			t.Logf("Names: %s, Id: %s, State: %s, NodeName: %s", row.Names[0], row.Id, row.State, row.Portainer.Agent.NodeName)
		}
	}

	// decodedBody, erro := res.DecodeBodySliceMap()
	// if erro != nil {
	// 	t.Error(erro)
	// }
	// // t.Logf("%#v", decodedBody)
	// t.Logf("len=%d", len(decodedBody))
	// // buf := make(map[string]map[string]string)
	// // for _, row := range decodedBody {
	// // 	buf[row["Names"]] = row
	// // 	// t.Logf("%#v", row)
	// // }
	// // t.Logf("len=%d", len(buf))
}
