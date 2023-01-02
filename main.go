package main

import (
	"fmt"
	"web-server-demo/pkg/config"
	"web-server-demo/pkg/router"
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

func main(){
	engine := gin.Default()
	gin.SetMode(gin.DebugMode)
	router.InitRouter(engine)
	err := engine.Run(fmt.Sprintf( "%s:%d",config.GetString(config.ServerHost),config.GetInt(config.ServerPort)))
	if err != nil {
		klog.Fatal((err))
		return
	}
}
