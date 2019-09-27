// query services
package services

import "github.com/tigershi/singletongo/service-facade/models"

type L3QueryService interface {
	GetComponentTranslation(req *models.ComponentMessageReq) (*models.ComponentQueryResult, error)
}
