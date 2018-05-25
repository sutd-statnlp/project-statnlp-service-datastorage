package config

import (
	"../api"
	"../core/factory"
	"../core/gateway"
	"../core/interactor"
	"github.com/gin-gonic/gin"
)

// CreateAddObjectAPI .
func CreateAddObjectAPI(factory factory.ObjectFactory, gateway gateway.ObjectGateway) api.AddObjectAPI {
	return api.AddObjectAPIImpl{
		Interactor: interactor.AddObjectInteractorImpl{
			Factory: factory,
			Gateway: gateway,
		},
	}
}

// CreateGetObjectAPI .
func CreateGetObjectAPI(gateway gateway.ObjectGateway) api.GetObjectAPI {
	return api.GetObjectAPIImpl{
		Interactor: interactor.GetObjectInteractorImpl{
			Gateway: gateway,
		},
	}
}

// SetAPIPaths sets API paths.
func SetAPIPaths(router *gin.Engine) {
	objectFactory := CreateObjectFactory()
	objectGateWay := CreateObjectGateway()
	addObjectAPI := CreateAddObjectAPI(objectFactory, objectGateWay)
	getObjectAPI := CreateGetObjectAPI(objectGateWay)

	apiGroup := router.Group("/api")
	{
		apiGroup.POST("/objects/:objectName", addObjectAPI.Add)
		apiGroup.GET("/objects/:objectName", getObjectAPI.Get)
		apiGroup.GET("/objects/:objectName/:objectId", getObjectAPI.GetByObjectID)
	}
}
