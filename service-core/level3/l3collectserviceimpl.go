// baseservice
package level3

import "github.com/tigershi/singletongo/service-facade/models"

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

func (qs *l3CollectServiceImpl) UpdateComponentTranslation(req models.ComponentCollectReq) (result models.ComponentCollectResult, err error) {

}

func (qs *l3CollectServiceImpl) DelComponentTranslation(req models.ComponentMessageReq) error {

}
