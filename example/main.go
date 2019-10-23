package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/tigershi/singletongo/example/bundle"
	_ "github.com/tigershi/singletongo/example/docs"
	"github.com/tigershi/singletongo/example/routers"
)

//@title Golang Translation API
//@version 1.0
//@description  Golang api of demo
//@termsOfService http://github.com

//@contact.name API Support
//@contact.url http://www.cnblogs.com
//@contact.email ×××@qq.com

//@host localhost:8000
func main() {
	gin.SetMode(gin.ReleaseMode)
	defaultRouter := gin.Default()
	router := routers.InitTransRouter(defaultRouter)
	url := ginSwagger.URL("http://localhost:8000/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.Run(":8000")

}
