package scheduler

import (
	"reflect"
	"testing"
)

type scheduleTestCase struct {
	name               string
	containers         Containers
	nodes              Nodes
	expectedCnt        int
	expectedPlacements []ContainerPlacement
}

func TestScheduler(t *testing.T) {
	testCases := []scheduleTestCase{
		{
			name: "singular_fit-less",
			containers: Containers{
				{
					Unit: Unit{
						UnitId: "xyz",
						Cpu:    2,
						Ram:    128,
						Constraints: map[string]any{
							"eu1":         nil,
							string(arm64): nil,
						},
					},
					Count: 1,
				}},
			nodes: Nodes{
				{
					Unit: Unit{UnitId: "node1",
						Cpu: 3,
						Ram: 256,
						Constraints: map[string]any{
							"eu1":         nil,
							string(arm64): nil,
						}},
				},
			},
			expectedCnt: 1,
			expectedPlacements: []ContainerPlacement{
				{
					NodeId: "node1",
					Containers: []Unit{
						{
							UnitId: "xyz",
							Cpu:    2,
							Ram:    128,
							Constraints: map[string]any{
								"eu1":         nil,
								string(arm64): nil,
							},
						},
					},
				},
			},
		},
		{
			name: "singular_fit-equal",
			containers: Containers{
				{
					Unit: Unit{UnitId: "xyz",
						Cpu: 2,
						Ram: 128,
						Constraints: map[string]any{
							"eu1":         nil,
							string(arm64): nil,
						}},
					Count: 1,
				}},
			nodes: Nodes{
				{
					Unit: Unit{UnitId: "node1",
						Cpu: 2,
						Ram: 128,
						Constraints: map[string]any{
							"eu1":         nil,
							string(arm64): nil,
						},
					},
				},
			},
			expectedCnt: 1,
			expectedPlacements: []ContainerPlacement{
				{
					NodeId: "node1",
					Containers: []Unit{
						{UnitId: "xyz",
							Cpu: 2,
							Ram: 128,
							Constraints: map[string]any{
								"eu1":         nil,
								string(arm64): nil,
							},
						}},
				},
			},
		},
		{
			name: "singular_fit-more",
			containers: Containers{
				{
					Unit: Unit{UnitId: "xyz",
						Cpu: 3,
						Ram: 256,
						Constraints: map[string]any{
							"eu1":         nil,
							string(arm64): nil,
						}},
					Count: 1,
				}},
			nodes: Nodes{
				{
					Unit: Unit{UnitId: "node1",
						Cpu: 2,
						Ram: 128,
						Constraints: map[string]any{
							"eu1":         nil,
							string(arm64): nil,
						}},
				},
			},
			expectedCnt:        0,
			expectedPlacements: []ContainerPlacement{},
		},
		{
			name: "multiple_fit-architecture,count",
			containers: Containers{
				{
					Unit: Unit{UnitId: "c2",
						Cpu: 2,
						Ram: 256,
						Constraints: map[string]any{
							"eu2":         nil,
							string(arm64): nil,
						}},
					Count: 2,
				},
			},
			nodes: Nodes{
				{
					Unit: Unit{UnitId: "node1",
						Cpu: 4,
						Ram: 512,
						Constraints: map[string]any{
							"eu2":          nil,
							string(x86_64): nil,
						}},
				},
				{
					Unit: Unit{UnitId: "node2",
						Cpu: 2,
						Ram: 256,
						Constraints: map[string]any{
							"eu2":         nil,
							string(arm64): nil,
						}},
				},
				{
					Unit: Unit{UnitId: "node3",
						Cpu: 2,
						Ram: 256,
						Constraints: map[string]any{
							"eu2":         nil,
							string(arm64): nil,
						}},
				},
			},
			expectedCnt: 2,
			expectedPlacements: []ContainerPlacement{
				{
					NodeId: "node2",
					Containers: []Unit{
						{
							UnitId: "c2",
							Cpu:    2,
							Ram:    256,
							Constraints: map[string]any{
								"eu2":         nil,
								string(arm64): nil,
							}},
					},
				},
				{
					NodeId: "node2",
					Containers: []Unit{
						{
							UnitId: "c2",
							Cpu:    2,
							Ram:    256,
							Constraints: map[string]any{
								"eu2":         nil,
								string(arm64): nil,
							}},
					},
				},
				{
					NodeId: "node3",
					Containers: []Unit{
						{
							UnitId: "c2",
							Cpu:    2,
							Ram:    256,
							Constraints: map[string]any{
								"eu2":         nil,
								string(arm64): nil,
							},
						},
					},
				},
				{
					NodeId: "node3",
					Containers: []Unit{
						{
							UnitId: "c2",
							Cpu:    2,
							Ram:    256,
							Constraints: map[string]any{
								"eu2":         nil,
								string(arm64): nil,
							},
						},
					},
				},
			},
		},
		{
			name: "multiple_fit-region",
			containers: Containers{
				{
					Unit: Unit{UnitId: "c1",
						Cpu: 2,
						Ram: 256,
						Constraints: map[string]any{
							"eu2":         nil,
							string(arm64): nil,
						}},
					Count: 2,
				},
				{
					Unit: Unit{UnitId: "c2",
						Cpu: 4,
						Ram: 512,
						Constraints: map[string]any{
							"eu1":         nil,
							string(arm64): nil,
						}},
					Count: 1,
				},
			},
			nodes: Nodes{
				{
					Unit: Unit{UnitId: "node1",
						Cpu: 4,
						Ram: 512,
						Constraints: map[string]any{
							"eu1":         nil,
							string(arm64): nil,
						}},
				},
				{
					Unit: Unit{UnitId: "node2",
						Cpu: 2,
						Ram: 256,
						Constraints: map[string]any{
							"eu2":         nil,
							string(arm64): nil,
						}},
				},
				{
					Unit: Unit{UnitId: "node3",
						Cpu: 2,
						Ram: 256,
						Constraints: map[string]any{
							"eu2":         nil,
							string(arm64): nil,
						}},
				},
			},
			expectedCnt: 3,
			expectedPlacements: []ContainerPlacement{
				{
					NodeId: "node1",
					Containers: []Unit{
						{
							UnitId: "c2",
							Cpu:    4,
							Ram:    512,
							Constraints: map[string]any{
								"eu1":         nil,
								string(arm64): nil,
							},
						},
					},
				},
				{
					NodeId: "node2",
					Containers: []Unit{
						{UnitId: "c1",
							Cpu: 2,
							Ram: 256,
							Constraints: map[string]any{
								"eu2":         nil,
								string(arm64): nil,
							},
						},
					},
				},
				{
					NodeId: "node3",
					Containers: []Unit{
						{
							UnitId: "c1",
							Cpu:    2,
							Ram:    256,
							Constraints: map[string]any{
								"eu2":         nil,
								string(arm64): nil,
							},
						},
					},
				},
			},
		},
		{
			name: "multiplefit_fit-priority",
			containers: Containers{
				{
					Unit: Unit{UnitId: "c1",
						Priority: 1,
						Cpu:      2,
						Ram:      128,
						Constraints: map[string]any{
							"eu1":         nil,
							string(arm64): nil,
						}},
					Count: 2,
				},
				{
					Unit: Unit{UnitId: "c1",
						Priority: 2,
						Cpu:      2,
						Ram:      256,
						Constraints: map[string]any{
							"eu1":         nil,
							string(arm64): nil,
						}},
					Count: 1,
				},
			},
			nodes: Nodes{
				{
					Unit: Unit{UnitId: "node1",
						Cpu: 4,
						Ram: 512,
						Constraints: map[string]any{
							"eu1":         nil,
							string(arm64): nil,
						}},
				},
				{
					Unit: Unit{UnitId: "node2",
						Cpu: 3,
						Ram: 256,
						Constraints: map[string]any{
							"eu1":         nil,
							string(arm64): nil,
						}},
				},
			},
			expectedCnt: 3,
			expectedPlacements: []ContainerPlacement{
				{
					NodeId: "node1",
					Containers: []Unit{
						{
							UnitId:   "c1",
							Priority: 1,
							Cpu:      2,
							Ram:      128,
							Constraints: map[string]any{
								"eu1":         nil,
								string(arm64): nil,
							}},
					},
				},
				{
					NodeId: "node1",
					Containers: []Unit{
						{
							UnitId:   "c1",
							Priority: 2,
							Cpu:      2,
							Ram:      256,
							Constraints: map[string]any{
								"eu1":         nil,
								string(arm64): nil,
							},
						},
					},
				},
				{
					NodeId: "node2",
					Containers: []Unit{
						{
							UnitId:   "c1",
							Priority: 1,
							Cpu:      2,
							Ram:      128,
							Constraints: map[string]any{
								"eu1":         nil,
								string(arm64): nil,
							},
						},
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
