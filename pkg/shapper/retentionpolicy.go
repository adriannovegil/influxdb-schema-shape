package shapper

import (
	"fmt"
)

// RetentionPolicy is a RetentionPolicy
type RetentionPolicy struct {
	Name               string
	Duration           string
	ShardGroupDuration string
	Replication        string
	Default            bool
}

// NewRetentionPolicy creates RetentionPolicies
func NewRetentionPolicy(args []interface{}) *RetentionPolicy {
	return &RetentionPolicy{
		Name:               iToS(args[0]),
		Duration:           iToS(args[1]),
		ShardGroupDuration: iToS(args[2]),
		Replication:        iToS(args[3]),
		Default:            args[4].(bool),
	}
}

func (rp *RetentionPolicy) String() string {
	return fmt.Sprintf(`  RP %v -> %v
    Default -> %v`, rp.Name, rp.Duration, rp.Default)
	// Default -> %v`, rp.Name, rp.Duration, rp.ShardGroupDuration, rp.Default)
}

func iToS(face interface{}) string {
	return fmt.Sprintf("%v", face)
}
