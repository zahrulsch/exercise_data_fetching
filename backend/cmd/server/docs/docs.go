// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/product": {
            "get": {
                "tags": [
                    "product"
                ],
                "summary": "api untuk mendapatkan data list product",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/database.FetchProductResult"
                        }
                    }
                }
            }
        },
        "/product/{id}": {
            "get": {
                "tags": [
                    "product"
                ],
                "summary": "api untuk mendapatkan data product satuan",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Product"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "database.FetchProductResult": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Product"
                    }
                }
            }
        },
        "model.Product": {
            "type": "object",
            "properties": {
                "catId": {
                    "type": "integer"
                },
                "countReview": {
                    "type": "integer"
                },
                "discountPercentage": {
                    "type": "integer"
                },
                "gaKey": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "imageUrl": {
                    "type": "string"
                },
                "imageUrlLarge": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "original_price": {
                    "type": "string"
                },
                "preorder": {
                    "type": "boolean"
                },
                "price": {
                    "type": "string"
                },
                "priceInt": {
                    "type": "integer"
                },
                "rating": {
                    "type": "number"
                },
                "url": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
