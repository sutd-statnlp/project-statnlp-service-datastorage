package db

import (
	"../core/entity"
	r "gopkg.in/gorethink/gorethink.v4"
)

// CommonDatabase .
type CommonDatabase interface {
	Connect() *r.Session
	CreateTable(tableName string) bool
	GetTableNames() []string
	ContainTable(tableName string) bool
	Insert(tableName string, instance entity.Object) entity.Object
	Find(tableName string) []entity.Object
	FindByID(tableName string, id string) entity.Object
	Update(tableName string, id string, instance entity.Object) entity.Object
}
