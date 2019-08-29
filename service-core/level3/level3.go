// leve3.go
package level3

import "github.com/tigershi/singletongo/service-facade/services"
import ""

func init() {
	initservices()
}

func initservices() {
	services.TranslationCollectService = GetL3CollectServiceInstance()
	services.TranslationQueryService = GetL3QueryServiceInstance()
}
func validatorinit() {

}
