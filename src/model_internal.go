package dash

import "time"

type CreateService struct {
	HostUrl        string
	ServiceUrl     string
	ServiceVersion string
	Config         serviceConfig
	// TODO service type (web, job, worker)
	Credentials registryCredentials
}

type UpdateServiceImage struct {
	ServiceId string
	// full url and path
	ImageUrl     string
	ImageVersion string
	Credentials  registryCredentials
}

type UpdateServiceConfig struct {
	ServiceId      string
	ServiceUrl     string
	ServiceVersion string
	Config         serviceConfig
	Credentials    registryCredentials
}

type ManageCreateService struct {
	WorkerId          string
	ServiceInstanceId string
	Service           CreateService
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

type ServiceLoad struct {
	ServiceId         string
	ServiceInstanceId string
	Cpu               int64
	Ram               int64
	ClientCreatedAt   time.Time
}

type serviceConfig struct {
	Regions        []string
	MinInstances   int64
	MaxInstances   int64
	CpuPerInstance int64
	RamPerInstance int64
}

type registryCredentials struct {
	Username string
	Password string
}
