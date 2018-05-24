package factory

import (
	"../generator"
)

// ObjectFactory .
type ObjectFactory interface {
	CreateWithID(instance interface{}) interface{}
	CreateWithIDAndTime(instance interface{}) interface{}
}

// ObjectFactoryImpl .
type ObjectFactoryImpl struct {
}

// CreateWithID create object instance with generated ID.
func (factory ObjectFactoryImpl) CreateWithID(instance interface{}) interface{} {
	return map[string]interface{}{
		"id":    generator.ObjectID(),
		"extra": instance,
	}
}

// CreateWithIDAndTime create object instance with generated ID, created and updated times.
func (factory ObjectFactoryImpl) CreateWithIDAndTime(instance interface{}) interface{} {
	return map[string]interface{}{
		"id":        generator.ObjectID(),
		"createdAt": generator.CurrentDateTime(),
		"updatedAt": generator.CurrentDateTime(),
		"extra":     instance,
	}
}
