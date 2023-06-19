// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "endalk200",
            "email": "eb808826@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/urls": {
            "get": {
                "description": "Get all short urls",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "urls"
                ],
                "summary": "Get all short urls",
                "responses": {
                    "200": {
                        "description": "short url created successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a short url",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "urls"
                ],
                "summary": "Create a short url",
                "parameters": [
                    {
                        "description": "Url Creation",
                        "name": "urlCreationRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.UrlCreationRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "short url created successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/urls/{shortUrl}": {
            "get": {
                "description": "Get a short url",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "urls"
                ],
                "summary": "Get a short url",
                "parameters": [
                    {
                        "type": "string",
                        "description": "short url",
                        "name": "shortUrl",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "short url created successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/{shortUrl}": {
            "get": {
                "description": "Redirect to initial url",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "urls"
                ],
                "summary": "Redirect to initial url",
                "parameters": [
                    {
                        "type": "string",
                        "description": "short url",
                        "name": "shortUrl",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "short url created successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.UrlCreationRequest": {
            "type": "object",
            "required": [
                "long_url",
                "user_id"
            ],
            "properties": {
                "long_url": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9808",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Blazingly Fast URL Shortener API",
	Description:      "Blazingly Fast URL Shortener API written in Go using Gin framework.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
