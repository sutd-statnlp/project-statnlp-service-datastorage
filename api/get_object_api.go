package api

import (
	"log"

	"../core/interactor"
	"github.com/gin-gonic/gin"
)

// GetObjectAPI .
type GetObjectAPI interface {
	GetAllObjectNames(context *gin.Context)
	Get(context *gin.Context)
	GetByObjectID(context *gin.Context)
}

// GetObjectAPIImpl .
type GetObjectAPIImpl struct {
	Interactor interactor.GetObjectInteractor
}

// GetAllObjectNames .
func (api GetObjectAPIImpl) GetAllObjectNames(context *gin.Context) {
	log.Println("request to get all object's names")
	body, err := api.Interactor.GetAll()
	if err != nil {
		context.JSON(400, err.Error())
		return
	}
	context.JSON(200, body)
}

// Get .
func (api GetObjectAPIImpl) Get(context *gin.Context) {
	objectName := context.Param("objectName")
	log.Println("request to get Object:", objectName)
	body := api.Interactor.Get(objectName)
	context.JSON(200, body)
}

// GetByObjectID .
func (api GetObjectAPIImpl) GetByObjectID(context *gin.Context) {
	objectName := context.Param("objectName")
	objectID := context.Param("objectId")
	log.Println("request to get Object:", objectName, " with ID:", objectID)
	body, err := api.Interactor.GetByObjectID(objectName, objectID)
	if err != nil {
		context.JSON(400, err.Error())
		return
	}
	context.JSON(200, body)
}
