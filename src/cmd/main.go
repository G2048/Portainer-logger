package main

import (
	"fmt"
	"log/slog"
	"portainer-logger/src/cmd/args"
	"portainer-logger/src/config"
	"portainer-logger/src/pkg/portainer"
	"strings"
)

// return another func for lazy getting result from api
func getContainers(client *portainer.PortainerApi) func() []portainer.ContainersResponse {
	return func() []portainer.ContainersResponse {
		res, rerr := client.Containers(1)
		if rerr != nil {
			panic(rerr)
		}

		var decodedBody []portainer.ContainersResponse
		_, err := res.DecodeBodyStruct(&decodedBody)
		if err != nil {
			panic(err)
		}

		return decodedBody
	}
}
func printContainersInfo(decodedBody []portainer.ContainersResponse) {
	for _, row := range decodedBody {
		if row.State == portainer.Running {
			fmt.Printf("Names: %s, Id: %s, State: %s, NodeName: %s\n", row.Names[0], row.Id, row.State, row.Portainer.Agent.NodeName)
		}
	}
}
func findContainerInfo(decodedBody []portainer.ContainersResponse, substring string) {
	for _, row := range decodedBody {
		if row.State == portainer.Running && strings.Contains(row.Names[0], substring) {
			fmt.Printf("Names: %s, Id: %s, State: %s, NodeName: %s\n", row.Names[0], row.Id, row.State, row.Portainer.Agent.NodeName)
		}
	}
}

func main() {
	config.InitLogger(slog.LevelInfo)
	config.InitLoadDotenv()

	settings := config.NewApiPortainerSettings()
	cmdArgs := args.NewCmdArgs()

	client := portainer.NewPortainerApi(settings.ApiKey)
	decodedBody := getContainers(client)
	switch {
	case cmdArgs.ContainersLogs.Is:
		findContainerInfo(decodedBody(), cmdArgs.Find)
	case cmdArgs.Find != "" && cmdArgs.ContainersInfo:
		findContainerInfo(decodedBody(), cmdArgs.Find)
	case cmdArgs.ContainersInfo:
		printContainersInfo(decodedBody())
	case cmdArgs.Find != "":
		findContainerInfo(decodedBody(), cmdArgs.Find)
	}
}
