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
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/todos": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GetTodos"
                ],
                "summary": "todo一覧を取得",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Todo"
                            }
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CreateTodo"
                ],
                "summary": "todoを作成",
                "parameters": [
                    {
                        "description": "Todo",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Todo"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.Todo"
                        }
                    }
                }
            }
        },
        "/todos/update": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UpdateTodo"
                ],
                "summary": "todoを更新",
                "parameters": [
                    {
                        "description": "Todo",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Todo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Todo"
                        }
                    }
                }
            }
        },
        "/todos/{id}": {
            "delete": {
                "tags": [
                    "DeleteTodo"
                ],
                "summary": "todoを削除",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "user一覧を取得する",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GetUsers"
                ],
                "summary": "user一覧を取得する",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.User"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "userを作成する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CreateUser"
                ],
                "summary": "userを作成する",
                "parameters": [
                    {
                        "description": "user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    }
                }
            }
        },
        "/users/auth0/{auth0_id}": {
            "get": {
                "description": "userをAuth0IDで取得する",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GetUserByAuth0ID"
                ],
                "summary": "userをAuth0IDで取得する",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Auth0ID",
                        "name": "auth0_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    }
                }
            }
        },
        "/users/delete/{id}": {
            "delete": {
                "description": "userを削除する",
                "tags": [
                    "DeleteUser"
                ],
                "summary": "userを削除する",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/update/{id}": {
            "post": {
                "description": "userを更新する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UpdateUser"
                ],
                "summary": "userを更新する",
                "parameters": [
                    {
                        "description": "user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/users/{id}": {
            "get": {
                "description": "userをIDで取得する",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GetUserByID"
                ],
                "summary": "userをIDで取得する",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Todo": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "entity.User": {
            "type": "object",
            "properties": {
                "auth0_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "description": "UUIDを想定",
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9090",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Recommend Swaggo API",
	Description:      "This is a recommend_swaggo server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
