// Code generated by swaggo/swag. DO NOT EDIT.

package openapi

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
        "/auth/me": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Get user by auth token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/sqlc.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/declarations": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "declarations"
                ],
                "summary": "Get all declarations",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 10,
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "organization_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "region_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.DeclarationGetAll"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/licenses": {
            "get": {
                "tags": [
                    "licenses"
                ],
                "summary": "Get all licenses",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "license_type",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.LicenseGetAll"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "tags": [
                    "auth"
                ],
                "summary": "login with email",
                "parameters": [
                    {
                        "description": "response.Success",
                        "name": "auth",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Success"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/organizations": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "organizations"
                ],
                "summary": "Get all organizations",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 10,
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "organization_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "region_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.OrganizationGetAll"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/payments": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "payments"
                ],
                "summary": "Get all payments",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 10,
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "organization_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.PaymentGetAll"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/types/licenses": {
            "get": {
                "tags": [
                    "types"
                ],
                "summary": "Get all license types",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/sqlc.LicenseType"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/types/payments": {
            "get": {
                "tags": [
                    "types"
                ],
                "summary": "Get all payment types",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/sqlc.PaymentType"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/types/statuses": {
            "get": {
                "tags": [
                    "types"
                ],
                "summary": "Get all statuses",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/sqlc.Status"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get users",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 10,
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "organization_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "region_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Users"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "response.Success",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UserCreateParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Success"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/sqlc.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update user by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "domain.UserCreateParams",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UserCreateParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Success"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete user by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Success"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.DeclarationGetAll": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "objects": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/sqlc.DeclarationGetAllRow"
                    }
                }
            }
        },
        "domain.LicenseGetAll": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "objects": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/sqlc.License"
                    }
                }
            }
        },
        "domain.Login": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "hashed_password": {
                    "type": "string"
                }
            }
        },
        "domain.OrganizationGetAll": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "objects": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/sqlc.Organization"
                    }
                }
            }
        },
        "domain.PaymentGetAll": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "objects": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/sqlc.Payment"
                    }
                }
            }
        },
        "domain.UserCreateParams": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "hashed_password": {
                    "type": "string"
                },
                "inn": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "region_id": {
                    "type": "integer"
                },
                "second_name": {
                    "type": "string"
                }
            }
        },
        "domain.Users": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "objects": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/sqlc.User"
                    }
                }
            }
        },
        "response.Error": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "response.Success": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "sqlc.DeclarationGetAllRow": {
            "type": "object",
            "properties": {
                "converage_of_the_danger_area": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "danger_rate": {
                    "type": "integer"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "life_insurance": {
                    "type": "string"
                },
                "location_info": {
                    "type": "string"
                },
                "organization_id": {
                    "type": "string"
                },
                "organization_name": {
                    "type": "string"
                },
                "proof": {
                    "type": "string"
                },
                "reasons_of_danger": {
                    "type": "string"
                },
                "residents_info": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "tech_document": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "sqlc.License": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "doc_file": {
                    "type": "string"
                },
                "document_number": {
                    "type": "string"
                },
                "granted_date": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "license_type": {
                    "type": "integer"
                },
                "lifetime": {
                    "type": "integer"
                },
                "organization_name": {
                    "type": "string"
                },
                "reestr_number": {
                    "type": "string"
                },
                "stir_number": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "work_category": {
                    "type": "string"
                }
            }
        },
        "sqlc.LicenseType": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "sqlc.Organization": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "parent_organization": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "sqlc.Payment": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "organization_id": {
                    "type": "string"
                },
                "requisites": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "type": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "sqlc.PaymentType": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "sqlc.Status": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "sqlc.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "hashed_password": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "inn": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "profile_picture": {
                    "type": "string"
                },
                "region_id": {
                    "type": "integer"
                },
                "role_id": {
                    "type": "integer"
                },
                "second_name": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Backend App",
	Description:      "This API contains the source for the backend app",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}