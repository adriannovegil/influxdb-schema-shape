package shapper

import (
	"fmt"

	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	client "github.com/influxdata/influxdb1-client/v2"
)

// Measurement is a measurement
type Measurement struct {
	Name   string
	Series int
	Tags   []*Tag
	Fields []*Field
}

// NewMeasurement creates measurements
func NewMeasurement(name string, db string, c client.Client) *Measurement {
	m := &Measurement{
		Name:   name,
		Tags:   make([]*Tag, 0),
		Fields: make([]*Field, 0),
	}
	m.getSeries(db, c)
	fmt.Println(m)
	m.getTags(db, c)
	m.getFields(db, c)
	return m
}

func (m *Measurement) getSeries(db string, c client.Client) {
	query := client.Query{
		Command:  fmt.Sprintf(`SHOW SERIES FROM "%v"`, m.Name),
		Database: db,
	}
	ret, err := c.Query(query)
	check(err)
	check(ret.Error())
	for _, val := range ret.Results[0].Series {
		m.Series = len(val.Values)
	}
}

func (m *Measurement) getTags(db string, c client.Client) {
	query := client.Query{
		Command:  fmt.Sprintf(`SHOW TAG KEYS FROM "%v"`, m.Name),
		Database: db,
	}
	ret, err := c.Query(query)
	check(err)
	check(ret.Error())
	for _, val := range ret.Results[0].Series {
		for _, tag := range val.Values {
			t := NewTag(tag[0].(string), db, m.Name, c)
			fmt.Println(t)
			m.Tags = append(m.Tags, t)
		}
	}
}

func (m *Measurement) getFields(db string, c client.Client) {
	query := client.Query{
		Command:  fmt.Sprintf(`SHOW FIELD KEYS FROM "%v"`, m.Name),
		Database: db,
	}
	ret, err := c.Query(query)
	check(err)
	check(ret.Error())
	for _, val := range ret.Results[0].Series {
		for _, field := range val.Values {
			f := NewField(field)
			fmt.Println(f)
			m.Fields = append(m.Fields, f)
		}
	}
}

func (m *Measurement) String() string {
	return fmt.Sprintf("  M %v -> %v", m.Name, m.Series)
}
