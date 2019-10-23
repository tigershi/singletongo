// query services
package services

import "github.com/tigershi/singletongo/example/models"

type L3QueryService interface {
	GetComponentTranslation(req *models.ComponentMessageReq) (*models.ComponentQueryResult, error)
}
