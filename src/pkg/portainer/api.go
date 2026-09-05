package portainer

import (
	"portainer-logger/src/pkg/api"
	"strconv"
)

const URL = "http://10.10.20.86:9000/api/endpoints"

type PortainerApi struct {
	*api.HttpApi
	ApiKey string
}

func NewPortainerApi(apiKey string) *PortainerApi {
	client := api.NewHttpApi(URL)
	client.AddHeader("X-API-Key", apiKey)
	return &PortainerApi{client, apiKey}
}
func (p *PortainerApi) baseAddHeaders() {
	p.AddHeader("X-API-Key:", p.ApiKey)
}

// Method PortainerApi for gettings info
// about Containers
func (p *PortainerApi) Containers(all int) (*api.Response, *api.ErrorHttp) {
	return p.Get("/1/docker/containers/json", api.QueryParams{"all": strconv.Itoa(all)})
}

// func(p *PortainerApi) () {
// 	return
// }
