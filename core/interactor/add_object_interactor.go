package interactor

import (
	"../entity"
	"../factory"
	"../gateway"
)

// AddObjectInteractor .
type AddObjectInteractor interface {
	Add(objectName string, instance entity.Object) entity.Object
}

// AddObjectInteractorImpl is the implementation of AddObjectInteractor interface.
type AddObjectInteractorImpl struct {
	Factory factory.ObjectFactory
	Gateway gateway.ObjectGateway
}

// Add inserts new object.
func (interactor AddObjectInteractorImpl) Add(objectName string, instance entity.Object) entity.Object {
	insertInstance := interactor.Factory.CreateWithIDAndTime(instance)
	return interactor.Gateway.Insert(objectName, insertInstance)
}
