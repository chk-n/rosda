package scheduler

func canSchedule(container, node Unit) bool {
	if container.Cpu > node.Cpu {
		return false
	} else if container.Ram > node.Ram {
		return false
	}

	for k := range container.Constraints {
		if _, ok := node.Constraints[k]; !ok {
			return false
		}
	}
	return true
}
