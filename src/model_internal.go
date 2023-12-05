package dash

import "time"

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
