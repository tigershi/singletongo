// leve3.go
package level3

import "github.com/tigershi/singletongo/service-facade/services"
import "github.com/tigershi/singletongo/service-facade/dao"

var componentDao *dao.ComponentDao

func init() {
	initservices()
}

func InitComponentDao(dao *dao.ComponentDao) {
	componentDao = dao
}

func initservices() {
	services.TranslationCollectService = GetL3CollectServiceInstance()
	services.TranslationQueryService = GetL3QueryServiceInstance()
}
func validatorinit() {

}
