package factory

import (
	"reflect"

	"../entity"
	"../generator"
)

// ObjectFactory .
type ObjectFactory interface {
	CreateWithID(instance entity.Object) entity.Object
	CreateWithIDAndTime(instance entity.Object) entity.Object
	UpdateWithTime(instance entity.Object) entity.Object
}

// ObjectFactoryImpl .
type ObjectFactoryImpl struct {
}

// CreateWithID create object instance with generated ID.
func (factory ObjectFactoryImpl) CreateWithID(instance entity.Object) entity.Object {
	return entity.Object{
		ID:    generator.ObjectID(),
		Extra: instance,
	}
}

// CreateWithIDAndTime create object instance with generated ID, created and updated times.
func (factory ObjectFactoryImpl) CreateWithIDAndTime(instance entity.Object) entity.Object {
	return entity.Object{
		ID:        generator.ObjectID(),
		CreatedAt: generator.CurrentDateTime(),
		UpdatedAt: generator.CurrentDateTime(),
		Extra:     instance,
	}
}

// UpdateWithTime update object instance with updated time.
func (factory ObjectFactoryImpl) UpdateWithTime(instance entity.Object) entity.Object {
	reflect.ValueOf(&instance).Elem().FieldByName("UpdatedAt").SetString(generator.CurrentDateTime())
	return instance
}
