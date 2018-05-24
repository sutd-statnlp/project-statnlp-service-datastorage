package config

import (
	"log"

	r "gopkg.in/gorethink/gorethink.v4"
)

// SetDbConnection sets db connection.
func SetDbConnection() bool {
	_, err := r.Connect(r.ConnectOpts{
		Address: "localhost:28015",
	})
	if err != nil {
		log.Fatalln(err)
	}
	return true
}
