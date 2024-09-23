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
                                "$ref": "#/definitions/usecase.TodoResponse"
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
                            "$ref": "#/definitions/usecase.CreateTodoRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/usecase.TodoResponse"
                        }
                    }
                }
            }
        },
        "/todos/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GetTodoByID"
                ],
                "summary": "todoを取得",
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
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Todo"
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
                "tags": [
                    "UpdateTodo"
                ],
                "summary": "todoを更新",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Todo",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usecase.UpdateTodoRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usecase.TodoResponse"
                        }
                    }
                }
            },
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
                                "$ref": "#/definitions/usecase.UserResponse"
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
                            "$ref": "#/definitions/usecase.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/usecase.UserResponse"
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
                            "$ref": "#/definitions/usecase.UserResponse"
                        }
                    }
                }
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
                            "$ref": "#/definitions/usecase.UserResponse"
                        }
                    }
                }
            },
            "put": {
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
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usecase.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
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
        "usecase.CreateTodoRequest": {
            "type": "object",
            "required": [
                "description",
                "title",
                "user_id"
            ],
            "properties": {
                "completed": {
                    "type": "boolean"
                },
                "description": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "usecase.CreateUserRequest": {
            "type": "object",
            "required": [
                "auth0_id",
                "email",
                "name",
                "role"
            ],
            "properties": {
                "auth0_id": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "role": {
                    "type": "string",
                    "enum": [
                        "admin",
                        "user"
                    ]
                }
            }
        },
        "usecase.TodoResponse": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean"
                },
                "createdAt": {
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
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "usecase.UpdateTodoRequest": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean"
                },
                "description": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "usecase.UpdateUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "usecase.UserResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "updatedAt": {
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
