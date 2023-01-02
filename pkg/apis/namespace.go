package apis

import (
	"net/http"
	"web-server-demo/pkg/service"
	"github.com/gin-gonic/gin"
)

func GetNamespaces(c *gin.Context){
	namespaces, err := service.GetNamespaces()
	if err != nil {
		c.JSON(http.StatusInternalServerError,err.Error())
	}
	c.JSON(http.StatusOK, namespaces)
}
