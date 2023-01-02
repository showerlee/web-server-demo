package apis

import (
	"net/http"
	"web-server-demo/pkg/service"
	"github.com/gin-gonic/gin"
)

func GetPods(c *gin.Context){
	namespaceName := c.Param("namespaceName")
	pods, err := service.GetPods(namespaceName)
	if err != nil {
		c.JSON(http.StatusInternalServerError,err.Error())
	}
	c.JSON(http.StatusOK, pods)
}
