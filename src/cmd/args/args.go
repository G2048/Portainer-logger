package args

import (
	"flag"
	"portainer-logger/src/pkg/portainer"
)

type ContainersLogs struct {
	Is          bool
	ContainerId string
	Node        string
	Tail        int
	Find        string
}
type CmdArgs struct {
	ContainersInfo bool
	Find           string
	Status         portainer.State
	ContainersLogs ContainersLogs
}

func NewCmdArgs() *CmdArgs {
	var containers = flag.Bool("containers", false, "Get containers from Portainer api")
	var find = flag.String("find", "", "Find container by mask")
	var status = flag.String("status", "all", `Get containers by status: "running", "exited", "created" and special "all"`)

	var isLogs = flag.Bool("logs", false, "Download logs from container")
	var containerIdLogs = flag.String("id", "", "Container id/name for logs")
	var nodeLogs = flag.String("node", "", "Container node name for logs")
	var tailLogs = flag.Int("tail", 10, "Print the last n record of logs; default tail=10")
	var findLogs = find

	flag.Parse()
	logs := ContainersLogs{
		*isLogs,
		*containerIdLogs,
		*nodeLogs,
		*tailLogs,
		*findLogs,
	}
	return &CmdArgs{
		ContainersInfo: *containers,
		Find:           *find,
		Status:         portainer.State(*status),
		ContainersLogs: logs,
	}
}
