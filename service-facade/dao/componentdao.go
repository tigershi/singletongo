// componentdao
package dao

import "github.com/tigershi/singletongo/service-facade/models"

type ComponentDao interface {
	GetComponentTranslation(productName string, version string, component string, language string) (*models.ComponentStore, error)
	UpdateComponentTranslation(compContent *models.ComponentStore) (*models.ComponentStore, error)
	DelComponentTranslation(productName string, version string, component string, language string) error
}
