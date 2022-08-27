package postgresql

type pgColumnConstraint struct {
	Name           string `json:"name"`
	ConstraintName string `json:"constraint_name"`
	ConstraintType string `json:"constraint"`
}
type constraints []pgColumnConstraint

func (cc constraints) isColumnUnique(name string) bool {
	for _, c := range cc {
		if c.Name == name && c.ConstraintType == "UNIQUE" {
			return true
		}
	}
	return false
}

func (cc constraints) isColumnPrimaryKey(name string) bool {
	for _, c := range cc {
		if c.Name == name && c.ConstraintType == "PRIMARY KEY" {
			return true
		}
	}
	return false
}

func (cc constraints) isColumnNotNull(name string) bool {
	for _, c := range cc {
		if c.Name == name && c.ConstraintType == "NOT NULL" {
			return true
		}
	}
	return false
}

func getPgColumnConstraintByName(constraints []pgColumnConstraint, name string) []pgColumnConstraint {
	var result []pgColumnConstraint
	for _, c := range constraints {
		if c.Name == name {
			result = append(result, c)
		}
	}
	return result
}
