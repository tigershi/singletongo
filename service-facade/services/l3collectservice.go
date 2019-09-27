// query services
package services

import "github.com/tigershi/singletongo/service-facade/models"

type L3CollectService interface {
	//create or update the component
	UpdateComponentTranslation(req *models.ComponentCollectReq) (result *models.ComponentCollectResult, err error)

	DelComponentTranslation(req *models.ComponentMessageReq) error
}
