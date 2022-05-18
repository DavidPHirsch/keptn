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
            "name": "Keptn Team",
            "url": "http://www.keptn.sh"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/scope": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get scopes",
                "tags": [
                    "Scopes"
                ],
                "summary": "Get scopes",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/model.GetScopesResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/secret": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get secrets",
                "tags": [
                    "Secrets"
                ],
                "summary": "Get secrets",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/model.GetSecretsResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
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
                "description": "Update an existing Secret",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Secrets"
                ],
                "summary": "Update a Secret",
                "parameters": [
                    {
                        "description": "The updated Secret",
                        "name": "secret",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Secret"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/model.Secret"
                        }
                    },
                    "400": {
                        "description": "Invalid payload",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
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
                "description": "Create a new Secret",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Secrets"
                ],
                "summary": "Create a Secret",
                "parameters": [
                    {
                        "description": "The new secret to be created",
                        "name": "secret",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Secret"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.Secret"
                        }
                    },
                    "400": {
                        "description": "Invalid payload",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
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
                "description": "Delete an existing Secret",
                "tags": [
                    "Secrets"
                ],
                "summary": "Delete a Secret",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The name of the secret",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The scope of the secret",
                        "name": "scope",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok"
                    },
                    "400": {
                        "description": "Invalid payload",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Data": {
            "type": "object",
            "additionalProperties": {
                "type": "string"
            }
        },
        "model.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "model.GetScopesResponse": {
            "type": "object",
            "properties": {
                "scopes": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.GetSecretResponseItem": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "keys": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "scope": {
                    "description": "Scope determines the scope of the secret (default=\"keptn-default\")",
                    "type": "string"
                }
            }
        },
        "model.GetSecretsResponse": {
            "type": "object",
            "properties": {
                "Secrets": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.GetSecretResponseItem"
                    }
                }
            }
        },
        "model.Secret": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "data": {
                    "$ref": "#/definitions/model.Data"
                },
                "name": {
                    "type": "string"
                },
                "scope": {
                    "description": "Scope determines the scope of the secret (default=\"keptn-default\")",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "x-token",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "develop",
	Host:             "",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Secret Service API",
	Description:      "This is the API documentation of the Secret Service.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
