package models

type L3BaseReq struct {
	Product string `json:"product, omitempty"`
	Version string `json:"version, omitempty"`
}

//messages requests
type L3QueryReq struct {
	Component string `json:"component, omitempty"`
	Language  string `json:"language, omitempty"`
	L3BaseReq
}

type ComponentMessageReq struct {
	L3QueryReq
}

//collect requests
type StringCollectReq struct {
	L3BaseReq
	Key   string `json:"key, omitempty"`
	Value string `json:"value, omitempty"`
}

type ComponentCollectReq struct {
	L3QueryReq
	Messages map[string]string `json:"messages, omitempty"`
}
