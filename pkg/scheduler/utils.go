package scheduler

func canSchedule(container Container, node Node) bool {
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

func CountResources(us []Unit) (cpu int64, ram int64) {
	for _, u := range us {
		cpu += int64(u.Cpu)
		ram += int64(u.Ram)
	}
	return cpu, ram
}
