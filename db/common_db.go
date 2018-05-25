package db

import r "gopkg.in/gorethink/gorethink.v4"

// CommonDatabase .
type CommonDatabase interface {
	Connect() *r.Session
	CreateTable(tableName string) bool
	GetTableNames() []string
	ContainTable(tableName string) bool
	Insert(tableName string, instance interface{}) interface{}
	Find(tableName string) []interface{}
	FindByID(tableName string, id string) interface{}
}
