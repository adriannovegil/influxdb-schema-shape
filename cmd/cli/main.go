package main

import (
	"flag"

	"devcircus<.com/schemashape/pkg/shapper"
)

var (
	host     *string
	username *string
	password *string
)

func init() {
	host = flag.String("host", "http://192.168.88.249:8086", "hostname of inlfux server")
	username = flag.String("u", "", "username for influx auth")
	password = flag.String("p", "", "password for influx auth")
	flag.Parse()
}

func main() {
	sc := shapper.NewSchamaShape(*host, *username, *password)
	sc.ShapeDatabases()
}
