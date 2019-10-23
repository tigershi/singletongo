package models

type HttpResponse struct {
	Code    int         `json:"code" example:"000"`
	Message string      `json:"message" example:"response infomation"`
	Data    interface{} `json:"data" `
}
