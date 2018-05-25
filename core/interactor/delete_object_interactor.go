package interactor

import (
	"../gateway"
)

// DeleteObjectInteractor .
type DeleteObjectInteractor interface {
	Delete(objectName string, objectID string) (bool, error)
}

// DeleteObjectInteractorImpl is the implementation of DeleteObjectInteractor interface.
type DeleteObjectInteractorImpl struct {
	Gateway gateway.ObjectGateway
}

// Delete delete object by ObjectID.
func (interactor DeleteObjectInteractorImpl) Delete(objectName string, objectID string) (bool, error) {
	return interactor.Gateway.Delete(objectName, objectID)
}
