package scheduler

type SystemArchitecture string

// A group of service instances (containers) or nodes
type Units []Unit

// Implements sort interface
func (u Units) Len() int {
	return len(u)
}

// Implements sort interface
func (u Units) Less(i, j int) bool {
	if u[i].Priority > u[j].Priority {
		return false
	}
	return u[i].GetResources() > u[j].GetResources()
}

// Implements sort interface
func (u Units) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func (u Units) GetResources() (c uint64) {
	for i := 0; i < len(u); i++ {
		c += u[i].GetResources()
	}
	return c
}

// A singular container or node
type Unit struct {
	UnitId   string
	Priority uint32
	// number of units
	Count       uint32
	Cpu         uint32
	Ram         uint32
	Constraints map[string]any
}

func (u *Unit) GetResources() uint64 {
	//! Count should not be multiplied here as its only used to schedule on different nodes
	return uint64(u.Cpu) * uint64(u.Ram)
}

type Placements struct {
	nTc ContainerPlacement // maps containers to node id
	cTn map[string]struct {
		pos    int
		nodeId string
	} // maps container id to position
	cs []string // list of container ids
}

func (p *Placements) Put(k string, v Unit) {
	p.cs = append(p.cs, k)
	p.nTc.Put(k, v)
}

func (p *Placements) removeLatest(cnt int) {
	// iteratively remove latest placed containers
	for i := 0; i < cnt; i++ {
		id := p.cs[len(p.cs)-1]
		p.removeContainer(id)
		p.cs = p.cs[:len(p.cs)-2]
	}
}

func (p *Placements) removeContainer(id string) {
	c, ok := p.cTn[id]
	if !ok {
		return
	}
	p.nTc.Remove(c.nodeId, c.pos)
}

func (p *Placements) RemoveById(k string, id string) {
	panic("implement me")
}

type ContainerPlacement map[string][]Unit

func (c ContainerPlacement) Put(nodeId string, container Unit) {
	if _, ok := c[nodeId]; !ok {
		c[nodeId] = []Unit{container}
		return
	}
	c[nodeId] = append(c[nodeId], container)
}

func (c ContainerPlacement) Remove(nodeId string, idx int) {
	c[nodeId] = append(c[nodeId][:idx], c[nodeId][idx+1:]...)
}

type NodeChange struct {
	Index int
	Cpu   uint32
	Ram   uint32
}
