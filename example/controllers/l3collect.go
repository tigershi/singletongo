// l3query
package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/tigershi/singletongo/example/level3"
	"github.com/tigershi/singletongo/example/models"
	"github.com/tigershi/singletongo/example/services"
)

var collectTrans services.L3CollectService = level3.GetL3CollectServiceInstance()

//@Summary create or update the component translation
//@Version 1.0
//@Accept  json
//@Produce  json
//@Param product path string true "product"
//@Param version path string true "version"
//@Param component path string true "component"
//@Param language path string true "language"
//@Param messages body object true "messages"
//@Success 200 object models.HttpResponse
//@Router /translation/singlecomponent/{product}/{version}/{component}/{language} [put]
func UpdateComponentTranslation(c *gin.Context) {

	result := models.HttpResponse{}

	productName := c.Param("product")
	version := c.Param("version")
	component := c.Param("component")
	language := c.Param("language")
	msgs := make(map[string]string)
	err1 := json.NewDecoder(c.Request.Body).Decode(&msgs)
	if err1 != nil {
		result.Code = 304
		result.Message = "parameter error"
		c.JSON(http.StatusOK, result)
		return
	}

	req := new(models.ComponentCollectReq)
	req.Product = productName
	req.Version = version
	req.Component = component
	req.Language = language
	req.Messages = msgs
	fmt.Printf("-----------------------------\n")
	log.WithFields(log.Fields{
		"product":   req.Product,
		"version":   req.Version,
		"component": req.Component,
		"language":  req.Language,
		"msg":       req.Messages["abc"],
	}).Info("-----------------------")
	respondata, err := collectTrans.UpdateComponentTranslation(req)
	if err != nil {
		result.Code = 304
		result.Message = "parameter error"
		c.JSON(http.StatusOK, result)
		return
	}

	result.Code = 200
	result.Message = "OK"
	result.Data = respondata
	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

//@Summary delete the component translation
//@Version 1.0
//@Param product path string true "product"
//@Param version path string true "version"
//@Param component path string true "component"
//@Param language path string true "language"
//@Success 200 object models.HttpResponse
//@Router /translation/singlecomponent/{product}/{version}/{component}/{language} [delete]
func DelComponentTranslation(c *gin.Context) {
	productName := c.Param("product")
	version := c.Param("version")
	component := c.Param("component")
	language := c.Param("language")
	req := new(models.ComponentMessageReq)
	req.Product = productName
	req.Version = version
	req.Component = component
	req.Language = language
	err := collectTrans.DelComponentTranslation(req)
	result := models.HttpResponse{}
	if err != nil {

		result.Code = 304
		result.Message = "delete component failure"
		c.JSON(http.StatusOK, result)
		return
	}

	result.Code = 200
	result.Message = "delete component successfully"
	c.JSON(http.StatusOK, result)
}
