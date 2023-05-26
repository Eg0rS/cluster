// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/ping": {
            "get": {
                "description": "ping-pong ops...",
                "responses": {}
            }
        },
        "/update/{user_id}": {
            "post": {
                "description": "Upsert user info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/profile.UpsertUserInfoReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/profile_responses.UpsertGoodResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/profile_responses.UpsertBadResponse"
                        }
                    },
                    "501": {
                        "description": "Not Implemented",
                        "schema": {
                            "$ref": "#/definitions/profile_responses.UpsertBadResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "profile.UpsertUserInfoReq": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer",
                    "example": 68
                },
                "city": {
                    "type": "string",
                    "example": "Москва"
                },
                "education": {
                    "type": "string",
                    "example": "Бакалавриат"
                },
                "firstname": {
                    "type": "string",
                    "example": "Владимир"
                },
                "patronymic": {
                    "type": "string",
                    "example": "Владимирович"
                },
                "surname": {
                    "type": "string",
                    "example": "Путин"
                },
                "university": {
                    "type": "string",
                    "example": "ИТМО"
                }
            }
        },
        "profile_responses.UpsertBadResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Bad request, error on service"
                }
            }
        },
        "profile_responses.UpsertGoodResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": ""
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swagger of API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
