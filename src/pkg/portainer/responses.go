package portainer

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"log/slog"
	"portainer-logger/src/pkg/api"
	"strings"
)

// Override base api response
type PortainerResponse struct {
	api.ResponseApi
}

// Override base DecodeBodyString() method for remove first 8 raw bytes
func (p PortainerResponse) DecodeBodyString() (string, error) {
	var builder strings.Builder
	scanner := bufio.NewScanner(bytes.NewReader(p.Body()))
	for scanner.Scan() {
		lineBytes := scanner.Bytes()
		if len(lineBytes) >= 8 {
			builder.Write(lineBytes[8:])
			builder.WriteByte('\n')
		}
	}
	return builder.String(), nil
}
func (p PortainerResponse) DecodeBodySliceMap() (jsonBody []api.JsonStringMap, err error) {
	panic(errors.New("Not implemented!"))
	err = json.Unmarshal(p.Body(), &jsonBody)
	if err != nil {
		slog.Error(err.Error())
	}
	return jsonBody, err
}
func (p PortainerResponse) DecodeBodyStruct(jsonBody any) (any, error) {
	panic(errors.New("Not implemented!"))
	err := json.Unmarshal(p.Body(), &jsonBody)
	if err != nil {
		slog.Error(err.Error())
	}
	return jsonBody, err
}

// type Containers struct {
type ContainersResponse struct {
	Command         string
	Created         int64
	HostConfig      HostConfig
	Id              string
	Image           string
	ImageID         string
	IsPortainer     bool
	Labels          map[string]string
	Mounts          []Mount
	Names           []string
	NetworkSettings map[string]NetworkSetting
	Portainer       Portainer
	Ports           []Port
	State           State
	Status          string
}
