package factory

import (
	"reflect"

	"../entity"
	"../generator"
)

// ObjectFactory .
type ObjectFactory interface {
	CreateWithDateTime(instance interface{}) entity.Object
	UpdateWithTime(updatedInstance entity.Object, instance interface{}) entity.Object
}

// ObjectFactoryImpl .
type ObjectFactoryImpl struct {
}

// CreateWithDateTime create object instance with created and updated times.
func (factory ObjectFactoryImpl) CreateWithDateTime(instance interface{}) entity.Object {
	return entity.Object{
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
