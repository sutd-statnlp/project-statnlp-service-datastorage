package config

import (
	"../core/gateway"
	"../db"
)

// CreateObjectGateway .
func CreateObjectGateway() gateway.ObjectGateway {
	var rethinkObjectGateway db.RethinkObjectGatewayImpl
	rethinkObjectGateway.Connect()
	return rethinkObjectGateway
}
