package db

import (
	"github.com/cockroachdb/pebble"
)

func NewConnection() (*pebble.DB, error) {
	conn, err := pebble.Open("_data", &pebble.Options{})
	if err != nil {
		return nil, err
	}

	return conn, nil
}
