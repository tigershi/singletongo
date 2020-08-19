// l3query
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tigershi/singletongo/example/configs"
	"github.com/tigershi/singletongo/example/level3"
	"github.com/tigershi/singletongo/example/models"
	"github.com/tigershi/singletongo/example/services"
)

var queryService services.L3QueryService = level3.GetL3QueryServiceInstance()

//@Summary get the component translation
//@Version 1.0
//@Produce  json
//@Param product path string true "product"
//@Param version path string true "version"
//@Param component path string true "component"
//@Param language path string true "language"
//@Success 200 object models.HttpResponse
//@Router /translation/singlecomponent/{product}/{version}/{component}/{language} [get]
func GetComponentTranslation(c *gin.Context) {
	configs.Logger.Info(c.Request.RequestURI)

	productName := c.Param("product")
	version := c.Param("version")
	component := c.Param("component")
	language := c.Param("language")

	req := new(models.ComponentMessageReq)
	req.Product = productName
	req.Version = version
	req.Component = component
	req.Language = language
	resultmodel, err := queryService.GetComponentTranslation(req)
	result := models.HttpResponse{}
	if err != nil {
		result.Code = 404
		result.Message = "no translation"
		c.JSON(http.StatusNotFound, result)
	}
	result.Code = 200
	result.Message = "OK"
	result.Data = resultmodel
	c.JSON(http.StatusOK, result)
	return

}
