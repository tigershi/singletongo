// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-10-22 16:18:50.3740401 +0800 CST m=+0.034982601

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://github.com",
        "contact": {
            "name": "API Support",
            "url": "http://www.cnblogs.com",
            "email": "×××@qq.com"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/translation/singlecomponent/{product}/{version}/{component}/{language}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "get the component translation",
                "parameters": [
                    {
                        "type": "string",
                        "description": "product",
                        "name": "product",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "version",
                        "name": "version",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "component",
                        "name": "component",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "language",
                        "name": "language",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.HttpResponse"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "create or update the component translation",
                "parameters": [
                    {
                        "type": "string",
                        "description": "product",
                        "name": "product",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "version",
                        "name": "version",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "component",
                        "name": "component",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "language",
                        "name": "language",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "messages",
                        "name": "messages",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.HttpResponse"
                        }
                    }
                }
            },
            "delete": {
                "summary": "delete the component translation",
                "parameters": [
                    {
                        "type": "string",
                        "description": "product",
                        "name": "product",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "version",
                        "name": "version",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "component",
                        "name": "component",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "language",
                        "name": "language",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.HttpResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.HttpResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 0
                },
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string",
                    "example": "response infomation"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8000",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Golang Translation API",
	Description: "Golang api of demo",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
