package api

import (
	"encoding/json"
	"log"

	"../core/interactor"
	"github.com/gin-gonic/gin"
)

// AddObjectAPI .
type AddObjectAPI interface {
	Add(context *gin.Context)
}

// AddObjectAPIImpl .
type AddObjectAPIImpl struct {
	Interactor interactor.AddObjectInteractor
}

// Add .
func (api AddObjectAPIImpl) Add(context *gin.Context) {
	objectName := context.Param("objectName")
	log.Println("request to add Object:", objectName)
	bytes, err := context.GetRawData()
	if err != nil {
		context.JSON(400, err.Error())
		return
	}
	var instance interface{}
	json.Unmarshal(bytes, &instance)
	body, err := api.Interactor.Add(objectName, instance)
	if err != nil {
		context.JSON(400, err.Error())
		return
	}
	context.JSON(200, body)
}
