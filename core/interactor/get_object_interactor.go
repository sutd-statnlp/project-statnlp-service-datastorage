package interactor

import (
	"../entity"
	"../gateway"
)

// GetObjectInteractor .
type GetObjectInteractor interface {
	Get(objectName string) []entity.Object
	GetByObjectID(objectName string, objectID string) (entity.Object, error)
	GetAll() ([]string, error)
}

// GetObjectInteractorImpl is the implementation of GetObjectInteractor interface.
type GetObjectInteractorImpl struct {
	Gateway gateway.ObjectGateway
}

// Get gets objects .
func (interactor GetObjectInteractorImpl) Get(objectName string) []entity.Object {
	return interactor.Gateway.Find(objectName)
}

// GetByObjectID gets object by ObjectID .
func (interactor GetObjectInteractorImpl) GetByObjectID(objectName string, objectID string) (entity.Object, error) {
	return interactor.Gateway.FindByID(objectName, objectID)
}

// GetAll gets all objects .
func (interactor GetObjectInteractorImpl) GetAll() ([]string, error) {
	return interactor.Gateway.FindAll()
}
