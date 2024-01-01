package rosda

type CreateServiceParams struct {
	ServiceId     string
	ImageUrl      string // NOTE idk if we will keep this field
	ImageVersion  string
	MinInstances  int64
	MaxInstances  int64
	Cpu           int64
	Ram           int64
	Datacenter    string
	CloudProvider string
}
type Service struct {
	ServiceId string
	// config
	MinInstances   int64
	MaxInstances   int64
	CpuPerInstance int64
	RamPerInstance int64
	Datacenters    []string
}

type ServiceInstance struct {
	InstanceId string
	Service
}

type Node struct {
	NodeId       string
	PublicKey    string
	Region       string
	MaxCpu       int64
	MaxRam       int64
	AvailableCpu int64
	AvailableRam int64
	IpAddress    string
	Port         string
}

type serviceConfig struct {
	Regions        []string
	MinInstances   int64
	MaxInstances   int64
	CpuPerInstance int64
	RamPerInstance int64
}
