// level3 project level3.go
package level3

import (
	"sync"

	"github.com/tigershi/singletongo/example/models"
	"github.com/tigershi/singletongo/example/services"
)

var l3QueryServiceInst services.L3QueryService
var l3QueryServiceOnce sync.Once

type l3QueryServiceImpl struct {
}

func GetL3QueryServiceInstance() services.L3QueryService {
	l3QueryServiceOnce.Do(func() {
		l3QueryServiceInst = new(l3QueryServiceImpl)
	})
	return l3QueryServiceInst
}

func (qs *l3QueryServiceImpl) GetComponentTranslation(req *models.ComponentMessageReq) (*models.ComponentQueryResult, error) {
	compStore, err := componentDao.GetComponentTranslation(req.Product, req.Version, req.Component, req.Language)
	if err != nil {
		return nil, err
	}

	result := new(models.ComponentQueryResult)
	result.L3BaseResult = compStore.L3BaseResult
	result.Messages = compStore.Messages
	return result, nil

}
