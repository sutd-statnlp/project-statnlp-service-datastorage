package interactor

import (
	"../entity"
	"../factory"
	"../gateway"
)

// UpdateObjectInteractor .
type UpdateObjectInteractor interface {
	Update(objectName string, objectID string, instance entity.Object) entity.Object
}

// UpdateObjectInteractorImpl is the implementation of UpdateObjectInteractor interface.
type UpdateObjectInteractorImpl struct {
	Factory factory.ObjectFactory
	Gateway gateway.ObjectGateway
}

// Update updates object.
func (interactor UpdateObjectInteractorImpl) Update(objectName string, objectID string, instance entity.Object) entity.Object {
	insertInstance := interactor.Factory.UpdateWithTime(instance)
	return interactor.Gateway.Update(objectName, objectID, insertInstance)
}
