package client

type Spec struct {
	// Used in API requests to filter only resources related to these team ids.
	// Used in the tables: ["escalation_policies", "incidents", "maintenance_windows", "services", "users"]
	TeamIds []string `yaml:"team_ids,omitempty" json:"team_ids"`
}
