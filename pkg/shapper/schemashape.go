package shapper

import (
	// "context"
	"fmt"
	"log"

	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	client "github.com/influxdata/influxdb1-client/v2"
)

// SchemaShape does things
type SchemaShape struct {
	Databases []*Database
	Client    client.Client
}

// NewSchamaShape returns things
func NewSchamaShape(host string, username string, password string) *SchemaShape {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     host,
		Username: username,
		Password: password,
	})
	check(err)
	return &SchemaShape{
		Databases: make([]*Database, 0),
		Client:    c,
	}
}

// ShapeDatabases shape the InfluxDB databases
func (sc *SchemaShape) ShapeDatabases() {
	query := client.Query{
		Command: "SHOW DATABASES",
	}
	ret, err := sc.Client.Query(query)
	check(err)
	check(ret.Error())

	for _, val := range ret.Results[0].Series[0].Values {
		db := NewDatabase(val[0].(string))
		sc.Databases = append(sc.Databases, db)
		fmt.Println(db)
		db.getRPs(sc.Client)
		db.getMeasurements(sc.Client)
		fmt.Println()
	}
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}
