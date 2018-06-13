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

// CreateUpdateObjectAPI .
func CreateUpdateObjectAPI(factory factory.ObjectFactory, gateway gateway.ObjectGateway) api.UpdateObjectAPI {
	return api.UpdateObjectAPIImpl{
		Interactor: interactor.UpdateObjectInteractorImpl{
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

// CreateDeleteObjectAPI .
func CreateDeleteObjectAPI(gateway gateway.ObjectGateway) api.DeleteObjectAPI {
	return api.DeleteObjectAPIImpl{
		Interactor: interactor.DeleteObjectInteractorImpl{
			Gateway: gateway,
		},
	}
}

// SetAPIPaths sets API paths.
func SetAPIPaths(router *gin.Engine, middleware gin.HandlerFunc) *gin.RouterGroup {
	objectFactory := CreateObjectFactory()
	objectGateWay := CreateObjectGateway()

	addObjectAPI := CreateAddObjectAPI(objectFactory, objectGateWay)
	getObjectAPI := CreateGetObjectAPI(objectGateWay)
	updateObjectAPI := CreateUpdateObjectAPI(objectFactory, objectGateWay)
	deleteObjectAPI := CreateDeleteObjectAPI(objectGateWay)

	apiGroup := router.Group("/api")
	apiGroup.Use(middleware)
	{
		apiGroup.POST("/objects/:objectName", addObjectAPI.Add)
		apiGroup.GET("/objects", getObjectAPI.GetAllObjectNames)
		apiGroup.GET("/objects/:objectName", getObjectAPI.Get)
		apiGroup.GET("/objects/:objectName/:objectId", getObjectAPI.GetByObjectID)
		apiGroup.PUT("/objects/:objectName/:objectId", updateObjectAPI.Update)
		apiGroup.DELETE("/objects/:objectName/:objectId", deleteObjectAPI.Delete)
	}
	return apiGroup
}
