// baseservice
package level3

import (
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tigershi/singletongo/example/models"
	"github.com/tigershi/singletongo/example/services"
)

type l3CollectServiceImpl struct {
}

var l3CollectServiceInst services.L3CollectService
var l3CollectServiceOnce sync.Once

func GetL3CollectServiceInstance() services.L3CollectService {
	l3CollectServiceOnce.Do(func() {
		l3CollectServiceInst = new(l3CollectServiceImpl)
	})
	return l3CollectServiceInst
}

func (qs *l3CollectServiceImpl) UpdateComponentTranslation(req *models.ComponentCollectReq) (*models.ComponentCollectResult, error) {
	store := new(models.ComponentStore)
	store.Product = req.Product
	store.Version = req.Version
	store.Component = req.Component
	store.Language = req.Language
	store.Messages = req.Messages
	store.Hash = time.Now().UnixNano()
	result, err := componentDao.UpdateComponentTranslation(store)
	if err != nil {
		return nil, err
	}

	log.WithFields(log.Fields{
		"product":   result.Product,
		"version":   result.Version,
		"component": result.Component,
		"language":  result.Language,
		"msg":       result.Messages["abc"],
	}).Info("-----------this the store component result------------")
	resultModel := new(models.ComponentCollectResult)
	resultModel.Product = result.Product
	resultModel.Version = result.Version
	resultModel.Component = result.Component
	resultModel.Language = result.Language
	resultModel.Messages = result.Messages
	return resultModel, nil
}

func (qs *l3CollectServiceImpl) DelComponentTranslation(req *models.ComponentMessageReq) error {
	return componentDao.DelComponentTranslation(req.Product, req.Version, req.Component, req.Language)
}
