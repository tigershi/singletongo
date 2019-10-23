package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tigershi/singletongo/example/controllers"
)

func InitTransRouter(router *gin.Engine) *gin.Engine {
	router.PUT("/translation/singlecomponent/:product/:version/:component/:language", controllers.UpdateComponentTranslation)
	router.DELETE("/translation/singlecomponent/:product/:version/:component/:language", controllers.DelComponentTranslation)
	router.GET("/translation/singlecomponent/:product/:version/:component/:language", controllers.GetComponentTranslation)
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	return router
}
