package scheduler

type SystemArchitecture string

// A singular container or node
type Unit struct {
	UnitId      string
	Priority    uint32
	Cpu         uint32
	Ram         uint32
	Constraints map[string]any
}

func (u *Unit) getResources() uint64 {
	//! Count should not be multiplied here as its only used to schedule on different nodes
	return uint64(u.Cpu) * uint64(u.Ram)
}

type Node struct {
	Unit
}

type Nodes []Node

// Implements sort interface
func (n Nodes) Len() int {
	return len(n)
}

// Implements sort interface
func (n Nodes) Less(i, j int) bool {
	if n[i].Priority > n[j].Priority {
		return false
	}
	return n[i].getResources() > n[j].getResources()
}

// Implements sort interface
func (n Nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n Nodes) getResources() (r uint64) {
	for i := 0; i < len(n); i++ {
		r += n[i].getResources()
	}
	return r
}

type Container struct {
	Unit
	// Number of units
	Count uint32
}

type Containers []Container

// Implements sort interface
func (c Containers) Len() int {
	return len(c)
}

// Implements sort interface
func (c Containers) Less(i, j int) bool {
	if c[i].Priority > c[j].Priority {
		return false
	}
	return c[i].getResources()*uint64(c[i].Count) > c[j].getResources()*uint64(c[j].Count)
}

// Implements sort interface
func (c Containers) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c Containers) getResources() (r uint64) {
	for i := 0; i < len(c); i++ {
		r += c[i].getResources()
	}
	return r
}

type placements struct {
	nTc map[string][]Unit // maps node id to containers
	cTn map[string]struct {
		pos    int
		nodeId string
	} // maps container id to position
	cs []string // list of container ids
}

func (p *placements) put(k string, v Unit) {
	p.cs = append(p.cs, k)

	if _, ok := p.nTc[k]; !ok {
		p.nTc[k] = []Unit{v}
		return
	}
	p.nTc[k] = append(p.nTc[k], v)
}

func (p *placements) removeLatest(cnt int) {
	// iteratively remove latest placed containers
	for i := 0; i < cnt; i++ {
		id := p.cs[len(p.cs)-1]
		p.removeContainer(id)
		p.cs = p.cs[:len(p.cs)-2]
	}
}

func (p *placements) removeContainer(id string) {
	c, ok := p.cTn[id]
	if !ok {
		return
	}

	p.nTc[id] = append(p.nTc[id][:c.pos], p.nTc[id][c.pos+1:]...)
}

// Returns cpu requirements of placed containers for a given nodeId. Returns 0 if nodeId not found
func (p *placements) CpuByNode(nodeId string) int {
	us, ok := p.nTc[nodeId]
	if !ok {
		return 0
	}
	cpu := 0
	for i := 0; i < len(us); i++ {
		cpu += int(us[i].Cpu)
	}
	return cpu

}

// Returns ram requirements of placed containers for a given nodeId. Returns 0 if nodeId not found
func (p *placements) RamByNode(nodeId string) int {
	us, ok := p.nTc[nodeId]
	if !ok {
		return 0
	}
	ram := 0
	for i := 0; i < len(us); i++ {
		ram += int(us[i].Ram)
	}
	return ram
}

// Returns a struct of placements
func (p *placements) Build() []ContainerPlacement {
	plc := make([]ContainerPlacement, 0, len(p.nTc))
	for k := range p.nTc {
		plc = append(plc, ContainerPlacement{
			NodeId:     k,
			Containers: p.nTc[k],
		})
	}
	return plc
}

type ContainerPlacement struct {
	NodeId     string
	Containers []Unit
}

type NodeChange struct {
	Index int
	Cpu   uint32
	Ram   uint32
}
