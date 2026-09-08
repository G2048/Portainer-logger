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
func printContainersInfo(decodedBody []portainer.ContainersResponse, status portainer.State) {
	for _, row := range decodedBody {
		if status == "all" || status == row.State {
			fmt.Printf("Names: %s, Id: %s, State: %s, Status: %s, NodeName: %s\n", row.Names[0], row.Id, row.State, row.Status, row.Portainer.Agent.NodeName)
			continue
		}
	}
}
func findContainerInfo(decodedBody []portainer.ContainersResponse, substring string, status portainer.State) (string, string) {
	var rowInfo string
	var rowId string
	var nodeName string

	for _, row := range decodedBody {
		if status == "all" || row.State == status {
			for _, name := range row.Names {
				if strings.Contains(name, substring) {
					rowInfo = fmt.Sprintf("Names: %s, Id: %s, State: %s, Status: %s, NodeName: %s", name, row.Id, row.State, row.Status, row.Portainer.Agent.NodeName)
					fmt.Println(rowInfo)
				}
			}
			rowId = row.Id
			nodeName = row.Portainer.Agent.NodeName
		}
	}
	return rowId, nodeName
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
	var ContainerId string
	var Node string

	config.InitLogger(slog.LevelInfo)
	config.InitLoadDotenv()

	settings := config.NewPortainerSettings()
	cmdArgs := args.NewCmdArgs()

	client := portainer.NewPortainerApi(settings.ApiKey)
	decodedBody := getContainers(client)
	switch {
	case cmdArgs.ContainersLogs.Is:
		if cmdArgs.ContainersLogs.Find != "" {
			ContainerId, Node = findContainerInfo(decodedBody(), cmdArgs.ContainersLogs.Find, cmdArgs.Status)
			if ContainerId == "" {
				panic(fmt.Errorf("ContainerId is empty!"))
			}
		} else if cmdArgs.ContainersLogs.ContainerId == "" {
			fmt.Println(fmt.Errorf("You must specify --container for --log flag!"))
			os.Exit(-1)
		} else if cmdArgs.ContainersLogs.Node == "" {
			fmt.Println(fmt.Errorf("You must specify --node for --log flag!"))
			os.Exit(-1)
		}
		body := getContainerLogs(
			client,
			ContainerId,
			Node,
			cmdArgs.ContainersLogs.Tail,
		)

		fmt.Printf("%s\n", body)
	case cmdArgs.Find != "" && cmdArgs.ContainersInfo:
		findContainerInfo(decodedBody(), cmdArgs.Find, cmdArgs.Status)
	case cmdArgs.ContainersInfo:
		printContainersInfo(decodedBody(), cmdArgs.Status)
	case cmdArgs.Find != "":
		findContainerInfo(decodedBody(), cmdArgs.Find, cmdArgs.Status)
	}
}
