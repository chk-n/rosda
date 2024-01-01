package scheduler

import "testing"

type canScheduleTestCase struct {
	name      string
	container Container
	node      Node
	expect    bool
}

func TestCanSchedule(t *testing.T) {
	testCases := []canScheduleTestCase{
		{
			name:      "fail_cpu",
			container: Container{Unit: Unit{Cpu: 10, Ram: 100, Constraints: map[string]any{}}},
			node:      Node{Unit: Unit{Cpu: 8, Ram: 1000, Constraints: map[string]any{}}},
			expect:    false,
		},
		{
			name:      "fail_ram",
			container: Container{Unit: Unit{Cpu: 2, Ram: 1000, Constraints: map[string]any{}}},
			node:      Node{Unit: Unit{Cpu: 4, Ram: 500, Constraints: map[string]any{}}},
			expect:    false,
		},
		{
			name:      "fail_constaint",
			container: Container{Unit: Unit{Cpu: 2, Ram: 100, Constraints: map[string]any{"eu1": nil}}},
			node:      Node{Unit: Unit{Cpu: 4, Ram: 1000, Constraints: map[string]any{}}},
			expect:    false,
		},
		{
			name:      "success",
			container: Container{Unit: Unit{Cpu: 2, Ram: 100, Constraints: map[string]any{"eu1": nil}}},
			node:      Node{Unit: Unit{Cpu: 4, Ram: 500, Constraints: map[string]any{"eu1": nil}}},
			expect:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := canSchedule(tc.container, tc.node)
			if got != tc.expect {
				t.Errorf("Test %s failed, expected %v, got %v", tc.name, tc.expect, got)
			}
		})
	}
}
