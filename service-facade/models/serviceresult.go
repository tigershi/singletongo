package models

type L3BaseResult struct {
	Product   string `json:"product, omitempty"`
	Version   string `json:"version, omitempty"`
	Component string `json:"component, omitempty"`
	Language  string `json:"language, omitempty"`
}

type StringCollectResult struct {
	L3BaseResult
	Key    string `json:"key, omitempty"`
	Status string `json:"status, omitempty"`
}

type ComponentQueryResult struct {
	L3BaseResult
	Messages map[string]string `json:"messages, omitempty"`
}

type ComponentCollectResult struct {
	ComponentQueryResult
}

type ComponentStore struct {
	ComponentQueryResult
	Hash uint64 `json:"hash, omitempty"`
}
