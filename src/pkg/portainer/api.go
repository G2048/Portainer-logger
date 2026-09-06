package portainer

import (
	"net/url"
	"portainer-logger/src/pkg/api"
	"portainer-logger/src/pkg/utils"
	"strconv"
)

const URL = "http://10.10.20.86:9000/api/endpoints"

type Container = string
type Node = string

// Convert PortainerResponse -> ResponseApi
func adapterResponse(r api.Response, err *api.ErrorHttp) (PortainerResponse, *api.ErrorHttp) {
	return PortainerResponse{r.(api.ResponseApi)}, err
}

type PortainerApi struct {
	*api.HttpApi
	ApiKey       string
	baseEndpoint string
}

func NewPortainerApi(apiKey string) *PortainerApi {
	client := api.NewHttpApi(URL)
	client.AddHeader("X-API-Key", apiKey)
	baseEndpoint := "/1/docker/containers/"
	return &PortainerApi{client, apiKey, baseEndpoint}
}
func (p *PortainerApi) baseAddHeaders() {
	p.AddHeader("X-API-Key:", p.ApiKey)
}
func (p PortainerApi) concatUrl(endpoints ...string) string {
	return utils.MustResult(url.JoinPath(p.baseEndpoint, endpoints...))
}

// Method PortainerApi for gettings info
// about Containers
func (p *PortainerApi) Containers(all int) (api.Response, *api.ErrorHttp) {
	return p.Get(p.concatUrl("json"), api.QueryParams{"all": strconv.Itoa(all)})
}

func (p *PortainerApi) ContainersLogs(container Container, nodeName Node, tail int) (PortainerResponse, *api.ErrorHttp) {
	p.AddHeader("X-PortainerAgent-Target", nodeName)
	return adapterResponse(p.Get(p.concatUrl(container, "/logs"), api.QueryParams{"stderr": "1", "stdout": "1", "tail": strconv.Itoa(tail)}))
}
