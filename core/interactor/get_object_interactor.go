package interactor

import (
	"../gateway"
)

// GetObjectInteractor .
type GetObjectInteractor interface {
	Get(objectName string) []interface{}
	GetByObjectID(objectName string, objectID string) interface{}
}

// GetObjectInteractorImpl is the implementation of GetObjectInteractor interface.
type GetObjectInteractorImpl struct {
	Gateway gateway.ObjectGateway
}

// Get gets objects .
func (interactor GetObjectInteractorImpl) Get(objectName string) []interface{} {
	return interactor.Gateway.Find(objectName)
}

// GetByObjectID gets object by ObjectID.
func (interactor GetObjectInteractorImpl) GetByObjectID(objectName string, objectID string) interface{} {
	return interactor.Gateway.FindByID(objectName, objectID)
}
