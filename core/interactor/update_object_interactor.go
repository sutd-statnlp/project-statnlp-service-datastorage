package interactor

import (
	"../entity"
	"../factory"
	"../gateway"
)

// UpdateObjectInteractor .
type UpdateObjectInteractor interface {
	Update(objectName string, objectID string, instance interface{}) (entity.Object, error)
}

// UpdateObjectInteractorImpl is the implementation of UpdateObjectInteractor interface.
type UpdateObjectInteractorImpl struct {
	Factory factory.ObjectFactory
	Gateway gateway.ObjectGateway
}

// Update updates object.
func (interactor UpdateObjectInteractorImpl) Update(objectName string, objectID string, instance interface{}) (entity.Object, error) {
	updateInstance, err := interactor.Gateway.FindByID(objectName, objectID)
	if err != nil {
		return entity.Object{}, err
	}
	updateInstance = interactor.Factory.UpdateWithTime(updateInstance, instance)
	return interactor.Gateway.Update(objectName, objectID, updateInstance)
}
