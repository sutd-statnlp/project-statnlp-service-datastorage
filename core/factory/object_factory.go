package factory

import (
	"reflect"

	"../entity"
	"../generator"
)

// ObjectFactory .
type ObjectFactory interface {
	CreateWithID(instance interface{}) entity.Object
	CreateWithIDAndTime(instance interface{}) entity.Object
	UpdateWithTime(updatedInstance entity.Object, instance interface{}) entity.Object
}

// ObjectFactoryImpl .
type ObjectFactoryImpl struct {
}

// CreateWithID create object instance with generated ID.
func (factory ObjectFactoryImpl) CreateWithID(instance interface{}) entity.Object {
	return entity.Object{
		ID:    generator.ObjectID(),
		Extra: instance,
	}
}

// CreateWithIDAndTime create object instance with generated ID, created and updated times.
func (factory ObjectFactoryImpl) CreateWithIDAndTime(instance interface{}) entity.Object {
	return entity.Object{
		ID:        generator.ObjectID(),
		CreatedAt: generator.CurrentDateTime(),
		UpdatedAt: generator.CurrentDateTime(),
		Extra:     instance,
	}
}

// UpdateWithTime update object instance with updated time.
func (factory ObjectFactoryImpl) UpdateWithTime(updatedInstance entity.Object, instance interface{}) entity.Object {
	reflect.ValueOf(&updatedInstance).Elem().FieldByName("UpdatedAt").SetString(generator.CurrentDateTime())
	reflect.ValueOf(&updatedInstance).Elem().FieldByName("Extra").Set(reflect.ValueOf(&instance))
	return updatedInstance
}
