// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "FDSAP Support"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/public/v1/dashboard/dashboardMenu": {
            "post": {
                "description": "Provide the menu to show in kPLUS",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DASHBOARD"
                ],
                "summary": "DASHBOARD-MENU",
                "parameters": [
                    {
                        "description": "Insti Code",
                        "name": "getInstiCode",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.InstiCodeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.DashboardMenuResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseModel"
                        }
                    }
                }
            }
        },
        "/public/v1/kplus/kplus_upon_open": {
            "get": {
                "description": "Provide the data that will be used by kPLUS upon opening the application",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "kPLUS"
                ],
                "summary": "kPLUS UPON OPEN",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.KplusUponOpenResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseModel"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.InstiCodeRequest": {
            "type": "object",
            "properties": {
                "insti_code": {
                    "type": "integer",
                    "example": 100
                }
            }
        },
        "response.DashboardMenuResponse": {
            "type": "object",
            "properties": {
                "dashboard": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "enable": {
                    "type": "integer"
                },
                "icon": {
                    "type": "integer"
                },
                "insti_code": {
                    "type": "integer"
                },
                "menu_item": {
                    "type": "string"
                },
                "order": {
                    "type": "integer"
                },
                "soon": {
                    "type": "integer"
                }
            }
        },
        "response.KplusUponOpenResponse": {
            "type": "object",
            "properties": {
                "action": {
                    "type": "string"
                },
                "created_by": {
                    "type": "integer"
                },
                "created_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image_url": {
                    "type": "string"
                },
                "last_updated_by": {
                    "type": "integer"
                },
                "last_updated_date": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "redirect_link": {
                    "type": "string"
                },
                "show": {
                    "type": "string"
                },
                "sub_message": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "response.ResponseModel": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "retCode": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/api/",
	Schemes:          []string{},
	Title:            "Fiber Example API",
	Description:      "FDSAP swagger template",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
