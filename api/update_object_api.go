package api

import (
	"encoding/json"
	"log"

	"../core/interactor"
	"github.com/gin-gonic/gin"
)

// UpdateObjectAPI .
type UpdateObjectAPI interface {
	Update(context *gin.Context)
}

// UpdateObjectAPIImpl .
type UpdateObjectAPIImpl struct {
	Interactor interactor.UpdateObjectInteractor
}

// Update .
func (api UpdateObjectAPIImpl) Update(context *gin.Context) {
	objectName := context.Param("objectName")
	objectID := context.Param("objectId")
	log.Println("request to update Object:", objectName, " with ID:", objectID)
	bytes, err := context.GetRawData()
	if err != nil {
		context.JSON(400, err.Error())
		return
	}
	var instance interface{}
	json.Unmarshal(bytes, &instance)
	body, err := api.Interactor.Update(objectName, objectID, instance)
	if err != nil {
		context.JSON(400, err.Error())
		return
	}
	context.JSON(200, body)
}
