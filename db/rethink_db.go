package db

import (
	"log"

	"../constant"
	r "gopkg.in/gorethink/gorethink.v4"
)

// RethinkDB .
type RethinkDB struct {
	Session *r.Session
}

// Connect .
func (db RethinkDB) Connect() *r.Session {
	session, err := r.Connect(r.ConnectOpts{
		Address: constant.RethinkHostAddress,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return session
}

// CreateTable .
func (db RethinkDB) CreateTable(tableName string) bool {
	_, err := r.DB(constant.DatabaseName).
		TableCreate(tableName).
		RunWrite(db.Session)
	if err != nil {
		log.Panic(err)
		return false
	}
	return true
}

// GetTableNames .
func (db RethinkDB) GetTableNames() []string {
	res, err := r.DB(constant.DatabaseName).
		TableList().
		Run(db.Session)
	if err != nil {
		log.Panic(err)
		return []string{}
	}
	defer res.Close()
	var rows []string
	err = res.All(&rows)
	if err != nil {
		log.Panic(err)
		return []string{}
	}
	return rows
}

// ContainTable .
func (db RethinkDB) ContainTable(tableName string) bool {
	res, err := r.DB(constant.DatabaseName).
		TableList().Contains(tableName).
		Run(db.Session)
	if err != nil {
		log.Panic(err)
		return false
	}
	defer res.Close()
	var row bool
	err = res.One(&row)
	if err != nil {
		log.Panic(err)
		return false
	}
	return row
}

// Insert .
func (db RethinkDB) Insert(tableName string, instance interface{}) interface{} {
	_, err := r.DB(constant.DatabaseName).
		Table(tableName).
		Insert(instance).
		RunWrite(db.Session)
	if err != nil {
		log.Panic(err)
		return []interface{}{}
	}
	return instance
}

// Find .
func (db RethinkDB) Find(tableName string) []interface{} {
	res, err := r.DB(constant.DatabaseName).
		Table(tableName).
		Run(db.Session)
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
func (db RethinkDB) FindByID(tableName string, id string) interface{} {
	res, err := r.DB(constant.DatabaseName).
		Table(tableName).
		Get(id).
		Run(db.Session)
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
