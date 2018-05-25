package api

import (
	"log"

	"../core/interactor"
	"github.com/gin-gonic/gin"
)

// DeleteObjectAPI .
type DeleteObjectAPI interface {
	Delete(context *gin.Context)
}

// DeleteObjectAPIImpl .
type DeleteObjectAPIImpl struct {
	Interactor interactor.DeleteObjectInteractor
}

// Delete .
func (api DeleteObjectAPIImpl) Delete(context *gin.Context) {
	objectName := context.Param("objectName")
	objectID := context.Param("objectId")
	log.Println("request to delete Object:", objectName, " with ID:", objectID)
	body, err := api.Interactor.Delete(objectName, objectID)
	if err != nil {
		context.JSON(400, err.Error())
		return
	}
	context.JSON(200, body)
}
