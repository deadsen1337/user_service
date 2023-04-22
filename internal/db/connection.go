package db

import (
	"github.com/cockroachdb/pebble"
	"log"
)

func NewConnection() *pebble.DB {
	conn, err := pebble.Open("_data", &pebble.Options{})
	if err != nil {
		log.Fatalf("could not open db connection: %s", err)
	}

	return conn
}
