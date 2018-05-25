package gateway

import "../entity"

// ObjectGateway .
type ObjectGateway interface {
	Insert(objectName string, instance entity.Object) entity.Object
	Find(objectName string) []entity.Object
	FindByID(objectName string, objectID string) entity.Object
	Update(objectName string, objectID string, instance entity.Object) entity.Object
}
