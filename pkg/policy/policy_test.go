package policy

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestHasChecks(t *testing.T) {
	filterTests := []struct {
		p      Policy
		result bool
	}{
		{
			result: true,
			p: Policy{
				Name: "aws",
				Policies: Policies{
					&Policy{
						Name: "test2",
						Policies: Policies{
							&Policy{
								Name: "test",
								Checks: []*Check{
									{
										Name: "Control-1",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			result: false,
			p: Policy{
				Name: "aws",
				Policies: Policies{
					&Policy{
						Name: "test2",
						Policies: Policies{
							&Policy{
								Name: "test",
							},
						},
					},
				},
			},
		}, {
			result: false,
			p: Policy{
				Name: "aws",
				Policies: Policies{
					&Policy{
						Name: "test",
					},
				},
			},
		},
		{
			p: Policy{
				Name: "aws",
				Policies: Policies{
					&Policy{
						Name: "level-1",
						Policies: Policies{
							&Policy{
								Name: "level-2",
								Policies: Policies{
									&Policy{
										Name: "level-3",
									},
								},
							},
						},
					},
				},
			},
			result: false,
		}, {
			p: Policy{
				Name: "aws",
				Policies: Policies{
					&Policy{
						Name: "level-1",
						Policies: Policies{
							&Policy{
								Name: "level-2",
								Policies: Policies{
									&Policy{
										Name: "level-3",
									},
								},
							},
						},
					},
				},
			},
			result: false,
		},
		{
			p: Policy{
				Name: "aws",
				Policies: Policies{
					&Policy{
						Name: "level-1",
						Policies: Policies{
							&Policy{
								Name: "level-2",
								Policies: Policies{
									&Policy{
										Name: "level-3",
										Checks: []*Check{
											{
												Name: "Control-1",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			result: true,
		},
		{
			p: Policy{
				Name: "aws",
				Policies: Policies{
					&Policy{
						Name: "level-1",
						Policies: Policies{
							&Policy{
								Name: "level-2",
								Policies: Policies{
									&Policy{
										Name: "level-3",
										Checks: []*Check{
											{
												Name: "Control-1",
											},
											{
												Name: "Control-2",
											},
											{
												Name: "Control-3",
											},
											{
												Name: "Control-4",
											},
											{
												Name: "Control-5",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			result: true,
		},
		{
			p: Policy{
				Name: "aws",
				Policies: Policies{
					&Policy{
						Name: "level-1",
						Policies: Policies{
							&Policy{
								Name: "level-2",
								Policies: Policies{
									&Policy{
										Name: "level-3",
										Checks: []*Check{
											{
												Name: "Control-1",
											},
											{
												Name: "Control-2",
											},
											{
												Name: "Control-3",
											},
											{
												Name: "Control-4",
											},
											{
												Name: "Control-5",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			result: true,
		},
	}
	for i, tt := range filterTests {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			diff := cmp.Diff(tt.result, tt.p.HasChecks(), cmpopts.IgnoreUnexported(Policy{}))
			if diff != "" {
				t.Fatalf("values are not the same %s", diff)
			}
		})
	}
}

func TestFilterPolicies(t *testing.T) {
	filterTests := []struct {
		p              Policy
		path           string
		expectError    bool
		expectedPolicy Policy
	}{
		{
			expectError: false,
			p: Policy{
				Name: "aws",
				Policies: Policies{
					&Policy{
						Name: "test2",
						Policies: Policies{
							&Policy{
								Name: "test",
							},
						},
					},
				},
			},
			path: "",
			expectedPolicy: Policy{
				Name: "aws",
				Policies: Policies{
					&Policy{
						Name: "test2",
						Policies: Policies{
							&Policy{
								Name: "test",
							},
						},
					},
				},
			},
		},
		{
			expectError: false,
			p: Policy{
				Name: "aws",
				Policies: Policies{
					&Policy{
						Name: "test2",
						Policies: Policies{
							&Policy{
								Name: "test",
							},
						},
					},
				},
			},
			path:           "test1",
			expectedPolicy: Policy{},
		}, {
			expectError: true,
			p: Policy{
				Name: "aws",
				Policies: Policies{
					&Policy{
						Name: "test",
					},
				},
			},
			path: "aws/test1",
			expectedPolicy: Policy{
				Name: "test",
			},
		},
		{
			p: Policy{
				Name: "aws",
				Policies: Policies{
					&Policy{
						Name: "level-1",
						Policies: Policies{
							&Policy{
								Name: "level-2",
								Policies: Policies{
									&Policy{
										Name: "level-3",
									},
								},
							},
						},
					},
				},
			},
			path: "level-1/level-2/level-3",
			expectedPolicy: Policy{
				Name: "level-3",
			},
		}, {
			p: Policy{
				Name: "aws",
				Policies: Policies{
					&Policy{
						Name: "level-1",
						Policies: Policies{
							&Policy{
								Name: "level-2",
								Policies: Policies{
									&Policy{
										Name: "level-3",
									},
								},
							},
						},
					},
				},
			},
			path: "level-1/level-2",
			expectedPolicy: Policy{
				Name: "level-2",
				Policies: Policies{
					&Policy{
						Name: "level-3",
					},
				},
			},
		},
		{
			p: Policy{
				Name: "aws",
				Policies: Policies{
					&Policy{
						Name: "level-1",
						Policies: Policies{
							&Policy{
								Name: "level-2",
								Policies: Policies{
									&Policy{
										Name: "level-3",
										Checks: []*Check{
											{
												Name: "Control-1",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			path: "level-1/level-2/level-3/Control-1",
			expectedPolicy: Policy{
				Name: "level-3",
				Checks: []*Check{
					{
						Name: "Control-1",
					},
				},
			},
		},
		{
			p: Policy{
				Name: "aws",
				Policies: Policies{
					&Policy{
						Name: "level-1",
						Policies: Policies{
							&Policy{
								Name: "level-2",
								Policies: Policies{
									&Policy{
										Name: "level-3",
										Checks: []*Check{
											{
												Name: "Control-1",
											},
											{
												Name: "Control-2",
											},
											{
												Name: "Control-3",
											},
											{
												Name: "Control-4",
											},
											{
												Name: "Control-5",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			path: "level-1/level-2/level-3",
			expectedPolicy: Policy{
				Name: "level-3",
				Checks: []*Check{
					{
						Name: "Control-1",
					},
					{
						Name: "Control-2",
					},
					{
						Name: "Control-3",
					},
					{
						Name: "Control-4",
					},
					{
						Name: "Control-5",
					},
				},
			},
		},
		{
			p: Policy{
				Name: "aws",
				Policies: Policies{
					&Policy{
						Name: "level-1",
						Policies: Policies{
							&Policy{
								Name: "level-2",
								Policies: Policies{
									&Policy{
										Name: "level-3",
										Checks: []*Check{
											{
												Name: "Control-1",
											},
											{
												Name: "Control-2",
											},
											{
												Name: "Control-3",
											},
											{
												Name: "Control-4",
											},
											{
												Name: "Control-5",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			path: "level-1/level-2",
			expectedPolicy: Policy{
				Name: "level-2",
				Policies: Policies{
					&Policy{
						Name: "level-3",
						Checks: []*Check{
							{
								Name: "Control-1",
							},
							{
								Name: "Control-2",
							},
							{
								Name: "Control-3",
							},
							{
								Name: "Control-4",
							},
							{
								Name: "Control-5",
							},
						},
					},
				},
			},
		},
	}
	for i, tt := range filterTests {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			diff := cmp.Diff(tt.expectedPolicy, tt.p.Filter(tt.path), cmpopts.IgnoreUnexported(Policy{}))
			if diff != "" && !tt.expectError {
				t.Fatalf("values are not the same %s", diff)
			}
		})
	}
}
