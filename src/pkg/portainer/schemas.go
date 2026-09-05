package portainer

type HostConfig struct {
	NetworkMode string
}
type Mount struct {
	Destination string
	Driver      string
	Mode        string
	Name        string
	Propagation string
	RW          bool
	Source      string
	Type        string
}
type NetworkSetting struct {
	Aliases             string
	DriverOpts          string
	EndpointID          int
	Gateway             int
	GlobalIPv6Address   int
	GlobalIPv6PrefixLen int
	IPAMConfig          map[string]string
	IPAddress           int
	IPPrefixLen         int
	IPv6Gateway         string
	Links               string
	MacAddress          string
	NetworkID           string
}
type NodeName struct {
	NodeName string
}

// type
type Portainer struct {
	Agent NodeName
}
type Port struct {
	IP          string
	PrivatePort int
	PublicPort  int
	Type        string
}

type State string

func (s State) Resolve() []State {
	return []State{"running"}
}

const Running State = "running"
const Exited State = "exited"
const Created State = "created"
