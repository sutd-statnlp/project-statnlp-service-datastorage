package api

import (
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
	var instance interface{}
	err := context.Bind(instance)
	if err != nil {
		context.JSON(400, err)
		return
	}
	body := api.Interactor.Add(objectName, instance)
	context.JSON(200, body)
}
