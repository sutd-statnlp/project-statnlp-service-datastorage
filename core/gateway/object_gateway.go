package gateway

// ObjectGateway .
type ObjectGateway interface {
	Insert(objectName string, instance interface{}) interface{}
	Find(objectName string) []interface{}
	FindByID(objectName string, objectID string) interface{}
}
