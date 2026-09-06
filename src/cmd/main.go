package main

import (
	"fmt"
	"log/slog"
	"os"
	"portainer-logger/src/cmd/args"
	"portainer-logger/src/config"
	"portainer-logger/src/pkg/portainer"
	"portainer-logger/src/pkg/utils"
	"strings"
)

// return another func for lazy getting result from api
func getContainers(client *portainer.PortainerApi) func() []portainer.ContainersResponse {
	return func() []portainer.ContainersResponse {
		res, erro := client.Containers(1)
		if erro != nil {
			panic(erro)
		}

		var decodedBody []portainer.ContainersResponse
		utils.MustResult(res.DecodeBodyStruct(&decodedBody))
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
func getContainerLogs(client *portainer.PortainerApi, container, node string, tail int) string {
	res, erro := client.ContainersLogs(container, node, tail)
	if erro != nil {
		panic(erro)
	}
	decodedBody := utils.MustResult(res.DecodeBodyString())
	return decodedBody
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
		if cmdArgs.ContainersLogs.ContainerId == "" {
			fmt.Println(fmt.Errorf("You must specify --container for --log flag!"))
			os.Exit(-1)
		} else if cmdArgs.ContainersLogs.Node == "" {
			fmt.Println(fmt.Errorf("You must specify --node for --log flag!"))
			os.Exit(-1)
		} else {
			body := getContainerLogs(
				client,
				cmdArgs.ContainersLogs.ContainerId,
				cmdArgs.ContainersLogs.Node,
				cmdArgs.ContainersLogs.Tail,
			)
			fmt.Printf("%s\n", body)
		}
	case cmdArgs.Find != "" && cmdArgs.ContainersInfo:
		findContainerInfo(decodedBody(), cmdArgs.Find)
	case cmdArgs.ContainersInfo:
		printContainersInfo(decodedBody())
	case cmdArgs.Find != "":
		findContainerInfo(decodedBody(), cmdArgs.Find)
	}
}
