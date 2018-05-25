package db

// ObjectGatewayImpl .
type ObjectGatewayImpl struct {
	Datatbase CommonDatabase
}

// Insert .
func (gateway ObjectGatewayImpl) Insert(objectName string, instance interface{}) interface{} {
	if !gateway.Datatbase.ContainTable(objectName) {
		gateway.Datatbase.CreateTable(objectName)
	}
	return gateway.Datatbase.Insert(objectName, instance)
}

// Find .
func (gateway ObjectGatewayImpl) Find(objectName string) []interface{} {
	return gateway.Datatbase.Find(objectName)
}

// FindByID .
func (gateway ObjectGatewayImpl) FindByID(objectName string, objectID string) interface{} {
	return gateway.Datatbase.FindByID(objectName, objectID)
}
