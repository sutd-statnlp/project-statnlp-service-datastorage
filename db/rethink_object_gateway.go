package db

import (
	"log"

	"../constant"
	r "gopkg.in/gorethink/gorethink.v4"
)

// RethinkObjectGatewayImpl .
type RethinkObjectGatewayImpl struct {
	Session *r.Session
}

// Connect .
func (gateway RethinkObjectGatewayImpl) Connect() bool {
	session, err := r.Connect(r.ConnectOpts{
		Address: constant.RethinkHostAddress,
	})
	if err != nil {
		log.Fatalln(err)
		return false
	}
	gateway.Session = session
	return true
}

// Insert .
func (gateway RethinkObjectGatewayImpl) Insert(objectName string, instance interface{}) interface{} {
	_, err := r.DB(constant.DatabaseName).Table(objectName).Insert(instance).RunWrite(gateway.Session)
	if err != nil {
		log.Panic(err)
		return []interface{}{}
	}
	return instance
}

// Find .
func (gateway RethinkObjectGatewayImpl) Find(objectName string) []interface{} {
	res, err := r.DB(constant.DatabaseName).Table(objectName).GetAll().Run(gateway.Session)
	if err != nil {
		log.Panic(err)
		return []interface{}{}
	}
	defer res.Close()
	var rows []interface{}
	err = res.All(&rows)
	if err != nil {
		log.Panic(err)
		return []interface{}{}
	}
	return rows
}

// FindByID .
func (gateway RethinkObjectGatewayImpl) FindByID(objectName string, objectID string) interface{} {
	res, err := r.DB(constant.DatabaseName).Table(objectName).Get(objectID).Run(gateway.Session)
	if err != nil {
		log.Panic(err)
		return []interface{}{}
	}
	defer res.Close()
	var row interface{}
	err = res.One(&row)
	if err != nil {
		log.Panic(err)
		return []interface{}{}
	}
	return row
}
