// baseservice
package level3

import (
	"time"

	"github.com/tigershi/singletongo/service-facade/models"
)

type l3CollectServiceImpl struct {
}

var l3CollectServiceInst *l3CollectServiceImpl
var l3CollectServiceOnce sync.Once

func GetL3CollectServiceInstance() *l3CollectServiceImpl {
	l3CollectServiceOnce.Do(func() {
		l3CollectServiceInst = &l3CollectServiceImpl{}
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
	store.Hash = time.Now().Nanosecond()
	result, err := componentDao.UpdateComponentTranslation(store)
	if err != nil {
		return nil, err
	}

	resultModel := &models.ComponentCollectResult{Product: result.Product, Version: result.Version, Component: result.Component, Language: result.Language, Messages: result.Messages}
	return resultModel, err
}

func (qs *l3CollectServiceImpl) DelComponentTranslation(req *models.ComponentMessageReq) error {
	return componentDao.DelComponentTranslation(req.Product, req.Version, req.Component, req.Language)
}
