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
        "/logout/": {
            "post": {
                "description": "to_register user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "responses": {}
            }
        },
        "/register": {
            "post": {
                "description": "to_register user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dal.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dal.OkRegisterResponse"
                        }
                    }
                }
            }
        },
        "/token": {
            "post": {
                "description": "to_auth user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.TokenRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.TokenResponse"
                        }
                    }
                }
            }
        },
        "/token-exist/": {
            "post": {
                "description": "to_auth user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.CheckTokenRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.CheckTokenResponse"
                        }
                    }
                }
            }
        },
        "/token/disable-login/": {
            "post": {
                "description": "to_register user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.DisableLoginRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/token/{userId}": {
            "delete": {
                "description": "delete refresh token by user id user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "dal.OkRegisterResponse": {
            "type": "object",
            "properties": {
                "AccessToken": {
                    "$ref": "#/definitions/dal.TokenResponse"
                },
                "RefreshToken": {
                    "$ref": "#/definitions/dal.TokenResponse"
                }
            }
        },
        "dal.TokenResponse": {
            "type": "object",
            "properties": {
                "TTL": {
                    "type": "integer"
                },
                "Token": {
                    "type": "string"
                }
            }
        },
        "dal.User": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "city": {
                    "type": "string"
                },
                "direction": {
                    "type": "string"
                },
                "education": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "eventDate": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "passwordHash": {
                    "type": "string"
                },
                "patronymic": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                },
                "university": {
                    "type": "string"
                }
            }
        },
        "handlers.CheckTokenRequest": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "handlers.CheckTokenResponse": {
            "type": "object",
            "properties": {
                "exists": {
                    "type": "boolean"
                }
            }
        },
        "handlers.DisableLoginRequest": {
            "type": "object",
            "properties": {
                "userId": {
                    "type": "string"
                }
            }
        },
        "handlers.GrantType": {
            "type": "string",
            "enum": [
                "password",
                "refresh_token"
            ],
            "x-enum-varnames": [
                "PasswordGrantType",
                "RefreshTokenGrantType"
            ]
        },
        "handlers.TokenRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "grant_type": {
                    "$ref": "#/definitions/handlers.GrantType"
                },
                "password": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "handlers.TokenResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "expires_in": {
                    "type": "integer"
                },
                "refresh_token": {
                    "type": "string"
                },
                "token_type": {
                    "type": "string"
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
