package api

import (
	"encoding/json"
	"log"

	"../core/entity"
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
	log.Println("request to add Object:", objectName, " with ID:", objectID)
	bytes, err := context.GetRawData()
	if err != nil {
		context.JSON(400, err.Error())
		return
	}
	var instance entity.Object
	json.Unmarshal(bytes, &instance)
	body := api.Interactor.Update(objectName, objectID, instance)
	context.JSON(200, body)
}
