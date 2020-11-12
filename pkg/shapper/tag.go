package shapper

import (
	"fmt"

	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	client "github.com/influxdata/influxdb1-client/v2"
)

// Tag is a tag
type Tag struct {
	Name        string
	Cardinality int
}

// NewTag creates tag
func NewTag(name string, db string, m string, c client.Client) *Tag {
	query := client.Query{
		Command:  fmt.Sprintf(`SHOW TAG VALUES FROM "%v" WITH KEY = "%v"`, m, name),
		Database: db,
	}
	ret, err := c.Query(query)
	check(err)
	check(ret.Error())
	t := &Tag{
		Name:        name,
		Cardinality: len(ret.Results[0].Series[0].Values),
	}
	return t
}

func (t *Tag) String() string {
	return fmt.Sprintf("    T %v -> %v", t.Name, t.Cardinality)
}
