// query services
package services

import "github.com/tigershi/singletongo/example/models"

type L3CollectService interface {
	//create or update the component
	UpdateComponentTranslation(req *models.ComponentCollectReq) (result *models.ComponentCollectResult, err error)
	//delete the component
	DelComponentTranslation(req *models.ComponentMessageReq) error
}
