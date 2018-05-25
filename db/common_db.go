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
	Insert(tableName string, instance interface{}) (entity.Object, error)
	Find(tableName string) []entity.Object
	FindByID(tableName string, id string) (entity.Object, error)
	Update(tableName string, id string, instance entity.Object) (entity.Object, error)
	Delete(objectName string, objectID string) (bool, error)
}
