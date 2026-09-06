package args

import (
	"flag"
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
	ContainersLogs ContainersLogs
}

func NewCmdArgs() *CmdArgs {
	var containers = flag.Bool("containers", false, "Get containers from Portainer api")
	var find = flag.String("find", "", "Find container by mask")
	var isLogs = flag.Bool("logs", false, "Download logs from container")
	var containerIdLogs = flag.String("container", "", "Container id/name for logs")
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
		ContainersLogs: logs,
	}
}
