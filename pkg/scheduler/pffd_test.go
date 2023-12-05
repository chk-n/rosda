package scheduler

import (
	"reflect"
	"testing"
)

type scheduleTestCase struct {
	name               string
	containers         Units
	nodes              Units
	expectedCnt        int
	expectedPlacements ContainerPlacement
}

func TestScheduler(t *testing.T) {
	testCases := []scheduleTestCase{
		{
			name: "singular_fit-less",
			containers: Units{
				{
					UnitId: "xyz",
					Cpu:    2,
					Ram:    128,
					Constraints: map[string]any{
						"eu1":         nil,
						string(arm64): nil,
					},
					Count: 1,
				}},
			nodes: Units{
				{
					UnitId: "node1",
					Cpu:    3,
					Ram:    256,
					Constraints: map[string]any{
						"eu1":         nil,
						string(arm64): nil,
					},
				},
			},
			expectedCnt: 1,
			expectedPlacements: map[string][]Unit{
				"node1": {
					{
						UnitId: "xyz",
						Cpu:    2,
						Ram:    128,
						Constraints: map[string]any{
							"eu1":         nil,
							string(arm64): nil,
						},
						Count: 1,
					}},
			},
		},
		{
			name: "singular_fit-equal",
			containers: Units{
				{
					UnitId: "xyz",
					Cpu:    2,
					Ram:    128,
					Constraints: map[string]any{
						"eu1":         nil,
						string(arm64): nil,
					},
					Count: 1,
				}},
			nodes: Units{
				{
					UnitId: "node1",
					Cpu:    2,
					Ram:    128,
					Constraints: map[string]any{
						"eu1":         nil,
						string(arm64): nil,
					},
				},
			},
			expectedCnt: 1,
			expectedPlacements: map[string][]Unit{
				"node1": {
					{
						UnitId: "xyz",
						Cpu:    2,
						Ram:    128,
						Constraints: map[string]any{
							"eu1":         nil,
							string(arm64): nil,
						},
						Count: 1,
					}},
			},
		},

		{
			name: "singular_fit-more",
			containers: Units{
				{
					UnitId: "xyz",
					Cpu:    3,
					Ram:    256,
					Constraints: map[string]any{
						"eu1":         nil,
						string(arm64): nil,
					},
					Count: 1,
				}},
			nodes: Units{
				{
					UnitId: "node1",
					Cpu:    2,
					Ram:    128,
					Constraints: map[string]any{
						"eu1":         nil,
						string(arm64): nil,
					},
				},
			},
			expectedCnt:        0,
			expectedPlacements: map[string][]Unit{},
		},
		{
			name: "multiple_fit-architecture,count",
			containers: Units{
				{
					UnitId: "c2",
					Cpu:    2,
					Ram:    256,
					Constraints: map[string]any{
						"eu2":         nil,
						string(arm64): nil,
					},
					Count: 2,
				},
			},
			nodes: Units{
				{
					UnitId: "node1",
					Cpu:    4,
					Ram:    512,
					Constraints: map[string]any{
						"eu2":          nil,
						string(x86_64): nil,
					},
				},
				{
					UnitId: "node2",
					Cpu:    2,
					Ram:    256,
					Constraints: map[string]any{
						"eu2":         nil,
						string(arm64): nil,
					},
				},
				{
					UnitId: "node3",
					Cpu:    2,
					Ram:    256,
					Constraints: map[string]any{
						"eu2":         nil,
						string(arm64): nil,
					},
				},
			},
			expectedCnt: 2,
			expectedPlacements: map[string][]Unit{
				"node2": {
					{
						UnitId: "c2",
						Cpu:    2,
						Ram:    256,
						Constraints: map[string]any{
							"eu2":         nil,
							string(arm64): nil,
						},
						Count: 2,
					},
				},
				"node3": {
					{
						UnitId: "c2",
						Cpu:    2,
						Ram:    256,
						Constraints: map[string]any{
							"eu2":         nil,
							string(arm64): nil,
						},
						Count: 2,
					},
				},
			},
		},
		{
			name: "multiple_fit-region",
			containers: Units{
				{
					UnitId: "c1",
					Cpu:    2,
					Ram:    256,
					Constraints: map[string]any{
						"eu2":         nil,
						string(arm64): nil,
					},
					Count: 2,
				},
				{
					UnitId: "c2",
					Cpu:    4,
					Ram:    512,
					Constraints: map[string]any{
						"eu1":         nil,
						string(arm64): nil,
					},
					Count: 1,
				},
			},
			nodes: Units{
				{
					UnitId: "node1",
					Cpu:    4,
					Ram:    512,
					Constraints: map[string]any{
						"eu1":         nil,
						string(arm64): nil,
					},
				},
				{
					UnitId: "node2",
					Cpu:    2,
					Ram:    256,
					Constraints: map[string]any{
						"eu2":         nil,
						string(arm64): nil,
					},
				},
				{
					UnitId: "node3",
					Cpu:    2,
					Ram:    256,
					Constraints: map[string]any{
						"eu2":         nil,
						string(arm64): nil,
					},
				},
			},
			expectedCnt: 3,
			expectedPlacements: map[string][]Unit{
				"node1": {
					{
						UnitId: "c2",
						Cpu:    4,
						Ram:    512,
						Constraints: map[string]any{
							"eu1":         nil,
							string(arm64): nil,
						},
						Count: 1,
					},
				},
				"node2": {
					{
						UnitId: "c1",
						Cpu:    2,
						Ram:    256,
						Constraints: map[string]any{
							"eu2":         nil,
							string(arm64): nil,
						},
						Count: 2,
					},
				},
				"node3": {
					{
						UnitId: "c1",
						Cpu:    2,
						Ram:    256,
						Constraints: map[string]any{
							"eu2":         nil,
							string(arm64): nil,
						},
						Count: 2,
					},
				},
			},
		},
		{
			name: "multiplefit_fit-priority",
			containers: Units{
				{
					UnitId:   "c1",
					Priority: 1,
					Cpu:      2,
					Ram:      128,
					Constraints: map[string]any{
						"eu1":         nil,
						string(arm64): nil,
					},
					Count: 2,
				},
				{
					UnitId:   "c1",
					Priority: 2,
					Cpu:      2,
					Ram:      256,
					Constraints: map[string]any{
						"eu1":         nil,
						string(arm64): nil,
					},
					Count: 1,
				},
			},
			nodes: Units{
				{
					UnitId: "node1",
					Cpu:    4,
					Ram:    512,
					Constraints: map[string]any{
						"eu1":         nil,
						string(arm64): nil,
					},
				},
				{
					UnitId: "node2",
					Cpu:    3,
					Ram:    256,
					Constraints: map[string]any{
						"eu1":         nil,
						string(arm64): nil,
					},
				},
			},
			expectedCnt: 3,
			expectedPlacements: map[string][]Unit{
				"node1": {
					{
						UnitId:   "c1",
						Priority: 1,
						Cpu:      2,
						Ram:      128,
						Constraints: map[string]any{
							"eu1":         nil,
							string(arm64): nil,
						},
						Count: 2,
					},
					{
						UnitId:   "c1",
						Priority: 2,
						Cpu:      2,
						Ram:      256,
						Constraints: map[string]any{
							"eu1":         nil,
							string(arm64): nil,
						},
						Count: 1,
					},
				},
				"node2": {
					{
						UnitId:   "c1",
						Priority: 1,
						Cpu:      2,
						Ram:      128,
						Constraints: map[string]any{
							"eu1":         nil,
							string(arm64): nil,
						},
						Count: 2,
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			scheduler := NewFirstFitDescending()
			gotPlacements, gotCnt := scheduler.Schedule(tc.containers, tc.nodes)
			if !reflect.DeepEqual(gotPlacements, tc.expectedPlacements) {
				t.Errorf("Test %s failed, expected %v, got %v", tc.name, tc.expectedPlacements, gotPlacements)
			}
			if tc.expectedCnt != gotCnt {
				t.Errorf("Test %s failed, expected %d, got %d", tc.name, tc.expectedCnt, gotCnt)
			}

		})
	}
}
