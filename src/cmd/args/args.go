package args

import "flag"

type CmdArgs struct {
	ContainersInfo bool
	Find           string
}

func NewCmdArgs() *CmdArgs {
	var containers = flag.Bool("containers", false, "Get containers from Portainer api")
	var find = flag.String("find", "", "Find container by mask")

	flag.Parse()
	return &CmdArgs{
		ContainersInfo: *containers,
		Find:           *find,
	}
}
