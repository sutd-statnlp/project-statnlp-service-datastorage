package db

import (
	"log"

	"../constant"
	"../core/entity"
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
		return false
	}
	return true
}

// GetTableNames .
func (db RethinkDB) GetTableNames() ([]string, error) {
	res, err := r.DB(constant.DatabaseName).
		TableList().
		Run(db.Session)
	if err != nil {
		return []string{}, err
	}
	defer res.Close()
	var rows []string
	err = res.All(&rows)
	if err != nil {
		return []string{}, err
	}
	return rows, nil
}

// ContainTable .
func (db RethinkDB) ContainTable(tableName string) bool {
	res, err := r.DB(constant.DatabaseName).
		TableList().Contains(tableName).
		Run(db.Session)
	if err != nil {

		return false
	}
	defer res.Close()
	var row bool
	err = res.One(&row)
	if err != nil {
		return false
	}
	return row
}

// Insert .
func (db RethinkDB) Insert(tableName string, instance interface{}) (entity.Object, error) {
	res, err := r.DB(constant.DatabaseName).
		Table(tableName).
		Insert(instance).
		RunWrite(db.Session)
	if err != nil {
		return entity.Object{}, err
	}
	result := instance.(entity.Object)
	result.ID = res.GeneratedKeys[0]
	return result, nil
}

// Find .
func (db RethinkDB) Find(tableName string) []entity.Object {
	res, err := r.DB(constant.DatabaseName).
		Table(tableName).
		Run(db.Session)
	if err != nil {
		return []entity.Object{}
	}
	defer res.Close()
	var rows []entity.Object
	err = res.All(&rows)
	if err != nil {
		return []entity.Object{}
	}
	return rows
}

// FindByID .
func (db RethinkDB) FindByID(tableName string, id string) (entity.Object, error) {
	res, err := r.DB(constant.DatabaseName).
		Table(tableName).
		Get(id).
		Run(db.Session)
	if err != nil {
		return entity.Object{}, err
	}
	defer res.Close()
	var row entity.Object
	err = res.One(&row)
	if err != nil {
		return entity.Object{}, err
	}
	return row, nil
}

// Update .
func (db RethinkDB) Update(tableName string, id string, instance entity.Object) (entity.Object, error) {
	_, err := r.DB(constant.DatabaseName).
		Table(tableName).
		Get(id).
		Update(instance).
		RunWrite(db.Session)
	if err != nil {
		return entity.Object{}, err
	}
	return instance, nil
}

// Delete .
func (db RethinkDB) Delete(tableName string, id string) (bool, error) {
	_, err := r.DB(constant.DatabaseName).
		Table(tableName).
		Get(id).
		Delete().
		RunWrite(db.Session)
	if err != nil {
		return false, err
	}
	return true, nil
}
