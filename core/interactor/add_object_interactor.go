package interactor

import (
	"../factory"
	"../gateway"
)

// AddObjectInteractor .
type AddObjectInteractor interface {
	Add(objectName string, instance interface{}) interface{}
}

// AddObjectInteractorImpl is the implementation of AddObjectInteractor interface.
type AddObjectInteractorImpl struct {
	Factory factory.ObjectFactory
	Gateway gateway.ObjectGateWay
}

// Add inserts new object.
func (interactor AddObjectInteractorImpl) Add(objectName string, instance interface{}) interface{} {
	insertInstance := interactor.Factory.CreateWithIDAndTime(instance)
	return interactor.Gateway.Insert(objectName, insertInstance)
}
