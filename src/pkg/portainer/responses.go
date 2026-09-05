package portainer

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
