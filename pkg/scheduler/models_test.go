package scheduler

import (
	"reflect"
	"sort"
	"testing"
)

func TestSortingContainers(t *testing.T) {
	testCases := []struct {
		name     string
		units    Containers
		expected Containers
	}{
		{
			name: "sort_containers-desc",
			units: Containers{
				{Count: 1, Unit: Unit{Cpu: 1, Ram: 100}},
				{Count: 1, Unit: Unit{Cpu: 3, Ram: 300}},
				{Count: 1, Unit: Unit{Cpu: 2, Ram: 200}},
			},
			expected: Containers{
				{Count: 1, Unit: Unit{Cpu: 3, Ram: 300}},
				{Count: 1, Unit: Unit{Cpu: 2, Ram: 200}},
				{Count: 1, Unit: Unit{Cpu: 1, Ram: 100}},
			},
		},
		{
			name: "sort_containers_with_count-desc",
			units: Containers{
				{Count: 2, Unit: Unit{Cpu: 4, Ram: 400}},
				{Unit: Unit{Cpu: 6, Ram: 600}},
				{Unit: Unit{Cpu: 5, Ram: 500}},
			},
			expected: Containers{
				{Count: 2, Unit: Unit{Cpu: 4, Ram: 400}},
				{Unit: Unit{Cpu: 6, Ram: 600}},
				{Unit: Unit{Cpu: 5, Ram: 500}},
			},
		},
		{
			name: "sort-priority",
			units: Containers{
				{Count: 1, Unit: Unit{Priority: 1, Cpu: 4, Ram: 400}},
				{Count: 1, Unit: Unit{Priority: 2, Cpu: 6, Ram: 600}},
				{Count: 1, Unit: Unit{Priority: 3, Cpu: 5, Ram: 500}},
			},
			expected: Containers{
				{Count: 1, Unit: Unit{Priority: 1, Cpu: 4, Ram: 400}},
				{Count: 1, Unit: Unit{Priority: 2, Cpu: 6, Ram: 600}},
				{Count: 1, Unit: Unit{Priority: 3, Cpu: 5, Ram: 500}},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sort.Sort(tc.units)

			if !reflect.DeepEqual(tc.units, tc.expected) {
				t.Errorf("Test %s failed, expected %v, got %v", tc.name, tc.expected, tc.units)
			}
		})
	}
}

func TestSortingNodes(t *testing.T) {
	testCases := []struct {
		name     string
		units    Nodes
		expected Nodes
	}{
		{
			name: "sort_nodes-desc",
			units: Nodes{
				{Unit{Cpu: 4, Ram: 400}},
				{Unit{Cpu: 6, Ram: 600}},
				{Unit{Cpu: 5, Ram: 500}},
			},
			expected: Nodes{
				{Unit{Cpu: 6, Ram: 600}},
				{Unit{Cpu: 5, Ram: 500}},
				{Unit{Cpu: 4, Ram: 400}},
			},
		},
		{
			name: "sort-priority",
			units: Nodes{
				{Unit{Priority: 1, Cpu: 4, Ram: 400}},
				{Unit{Priority: 2, Cpu: 6, Ram: 600}},
				{Unit{Priority: 3, Cpu: 5, Ram: 500}},
			},
			expected: Nodes{
				{Unit{Priority: 1, Cpu: 4, Ram: 400}},
				{Unit{Priority: 2, Cpu: 6, Ram: 600}},
				{Unit{Priority: 3, Cpu: 5, Ram: 500}},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sort.Sort(tc.units)

			if !reflect.DeepEqual(tc.units, tc.expected) {
				t.Errorf("Test %s failed, expected %v, got %v", tc.name, tc.expected, tc.units)
			}
		})
	}
}
