// validator project validator.go
package validator

type L3Validator struct {
	Validate(req L3BaseReq) bool
}

var L3Validators []L3Validator
