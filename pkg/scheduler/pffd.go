package scheduler

import (
	"sort"
)

// Implementation of a priority-based first fit descending algorithm
// with the goal to maximise the number of items put into bins i.e. number of containers scheduled onto nodes
// We use sorting:
// - containers by capacity (desc)
// - nodes by priority and resource availability (desc) and split the containers up by region and system architecture

type FirstFitDescending struct {
	ps  placements
	cnt int
}

func NewFirstFitDescending() *FirstFitDescending {
	return &FirstFitDescending{
		ps: placements{
			nTc: make(map[string][]Unit),
		},
	}
}

// Returns which containers were placed where and number of scheduled containers
func (b *FirstFitDescending) Schedule(containers Containers, nodes Nodes) ([]ContainerPlacement, int) {
	sort.Sort(containers)
	sort.Sort(nodes)

	for i := 0; i < len(containers); i++ {
		count := containers[i].Count
		for j := 0; j < len(nodes); j++ {
			// due to sorting and removing full nodes if this is true then there is no more capacity
			//if containers[i].GetResources() > nodes[j].GetResources() {
			//	return b.ps.nTc, b.cnt
			//}
			// if multiple instances schedule them on different nodes
			if count > 1 {
				// check if enough nodes available
				if len(nodes)-j < int(count) {
					// not enough nodes to accommodate multi-unit
					break
				}
				jSave := j

				var nChanges []NodeChange
				// TODO: add ability for some nodes to be hosted on same node (if user defines that)
				for c := 0; c < int(count); {
					if canSchedule(containers[i], nodes[j]) {
						// update node capacity
						nChanges = append(nChanges, NodeChange{
							Index: j,
							Cpu:   nodes[j].Cpu - containers[i].Cpu,
							Ram:   nodes[j].Ram - containers[i].Ram,
						})
						b.place(containers[i], nodes[j].UnitId)
						c++ // only up count if schedule was successful
					} else {
						j = jSave
						// no more capacity so we need to remove placed nodes
						b.removeLatest(c)
						break
					}
					j++
				}
				// apply changes
				for n := 0; n < len(nChanges); n++ {
					nodes[nChanges[n].Index].Cpu = nChanges[n].Cpu
					nodes[nChanges[n].Index].Ram = nChanges[n].Ram
				}
			} else {
				if canSchedule(containers[i], nodes[j]) {
					// update node capacity
					nodes[j].Cpu = nodes[j].Cpu - containers[i].Cpu
					nodes[j].Ram = nodes[j].Ram - containers[i].Ram
					b.place(containers[i], nodes[j].UnitId)
				}
			}

		}
		// sort nodes as node's capacity was updated
		sort.Sort(nodes)
	}
	return b.ps.Build(), b.cnt

}

func (b *FirstFitDescending) place(container Container, nodeId string) {
	b.cnt++
	b.ps.put(nodeId, container.Unit)
}

func (b *FirstFitDescending) removeLatest(cnt int) {
	b.cnt = b.cnt - cnt
	b.ps.removeLatest(cnt)
}
