package shapper

import (
	"fmt"

	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	client "github.com/influxdata/influxdb1-client/v2"
)

// Database is the abstraction
type Database struct {
	Name              string
	RetentionPolicies []*RetentionPolicy
	Measurements      []*Measurement
	Series            []*Serie
}

// NewDatabase returns a Database
func NewDatabase(name string) *Database {
	return &Database{Name: name}
}

func (db *Database) String() string {
	return fmt.Sprintf("DB %v", db.Name)
}

func (db *Database) getRPs(c client.Client) {
	query := client.Query{
		Command:  fmt.Sprintf("SHOW RETENTION POLICIES ON %v", db.Name),
		Database: db.Name,
	}
	ret, err := c.Query(query)
	check(err)
	check(ret.Error())
	for _, val := range ret.Results[0].Series {
		rp := NewRetentionPolicy(val.Values[0])
		fmt.Println(rp)
		db.RetentionPolicies = append(db.RetentionPolicies, rp)
	}
}

func (db *Database) getMeasurements(c client.Client) {
	query := client.Query{
		Command:  "SHOW MEASUREMENTS",
		Database: db.Name,
	}
	ret, err := c.Query(query)
	check(err)
	check(ret.Error())
	for _, val := range ret.Results[0].Series {
		for _, meas := range val.Values {
			m := NewMeasurement(meas[0].(string), db.Name, c)
			db.Measurements = append(db.Measurements, m)
		}
	}
}
