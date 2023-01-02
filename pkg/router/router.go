package router

import (
	"web-server-demo/pkg/apis"
	"web-server-demo/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine){
	middleware.InitMiddleware(r)
	r.GET("/ping",apis.Ping)
	r.GET("/namespaces",apis.GetNamespaces)
	r.GET("/namespace/:namespaceName/pods",apis.GetPods)
}
 