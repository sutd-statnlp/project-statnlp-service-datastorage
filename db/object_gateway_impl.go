package db

import "../core/entity"

// ObjectGatewayImpl .
type ObjectGatewayImpl struct {
	Datatbase CommonDatabase
}

// Insert .
func (gateway ObjectGatewayImpl) Insert(objectName string, instance entity.Object) entity.Object {
	if !gateway.Datatbase.ContainTable(objectName) {
		gateway.Datatbase.CreateTable(objectName)
	}
	return gateway.Datatbase.Insert(objectName, instance)
}

// Find .
func (gateway ObjectGatewayImpl) Find(objectName string) []entity.Object {
	return gateway.Datatbase.Find(objectName)
}

// FindByID .
func (gateway ObjectGatewayImpl) FindByID(objectName string, objectID string) entity.Object {
	return gateway.Datatbase.FindByID(objectName, objectID)
}

// Update .
func (gateway ObjectGatewayImpl) Update(objectName string, objectID string, instance entity.Object) entity.Object {
	return gateway.Datatbase.Update(objectName, objectID, instance)
}
