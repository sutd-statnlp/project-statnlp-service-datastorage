package config

import (
	"../core/gateway"
	"../db"
)

// CreateObjectGateway .
func CreateObjectGateway() gateway.ObjectGateway {
	var database db.RethinkDB
	database.Session = database.Connect()
	rethinkObjectGateway := db.ObjectGatewayImpl{
		Datatbase: database,
	}
	return rethinkObjectGateway
}
