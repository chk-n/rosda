package scheduler

import (
	"reflect"
	"sort"
	"testing"
)

type sortingTestCase struct {
	name     string
	units    Units
	expected Units
}

func TestSorting(t *testing.T) {
	testCases := []sortingTestCase{
		{
			name: "sort_containers-desc",
			units: Units{
				{Cpu: 1, Ram: 100},
				{Cpu: 3, Ram: 300},
				{Cpu: 2, Ram: 200},
			},
			expected: Units{
				{Cpu: 3, Ram: 300},
				{Cpu: 2, Ram: 200},
				{Cpu: 1, Ram: 100},
			},
		},
		{
			name: "sort_nodes-desc",
			units: Units{
				{Cpu: 4, Ram: 400},
				{Cpu: 6, Ram: 600},
				{Cpu: 5, Ram: 500},
			},
			expected: Units{
				{Cpu: 6, Ram: 600},
				{Cpu: 5, Ram: 500},
				{Cpu: 4, Ram: 400},
			},
		},
		{
			name: "sort-priority",
			units: Units{
				{Priority: 1, Cpu: 4, Ram: 400},
				{Priority: 2, Cpu: 6, Ram: 600},
				{Priority: 3, Cpu: 5, Ram: 500},
			},
			expected: Units{
				{Priority: 1, Cpu: 4, Ram: 400},
				{Priority: 2, Cpu: 6, Ram: 600},
				{Priority: 3, Cpu: 5, Ram: 500},
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
