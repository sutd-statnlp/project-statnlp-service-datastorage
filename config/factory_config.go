package config

import (
	"../core/factory"
)

// CreateObjectFactory .
func CreateObjectFactory() factory.ObjectFactory {
	return factory.ObjectFactoryImpl{}
}
