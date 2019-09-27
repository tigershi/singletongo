// validator project validator.go
package validator

import "github.com/tigershi/singletongo/service-facade/models"

type L3Validator interface {
	Validate(req L3BaseReq) bool
}
