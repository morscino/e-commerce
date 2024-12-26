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
        "/auth": {
            "post": {
                "description": "Creates a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create new user",
                "parameters": [
                    {
                        "description": "data to sign up new user",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SignUpDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseObject"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Logs in a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Login  user",
                "parameters": [
                    {
                        "description": "data to log in a user",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SignInDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseObject"
                        }
                    }
                }
            }
        },
        "/orders": {
            "get": {
                "description": "Gets All Orders",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Get All Orders",
                "parameters": [
                    {
                        "description": "data to query for all ",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.APIPagingDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "post": {
                "description": "Places a new Order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Place Order",
                "parameters": [
                    {
                        "description": "data to place new order",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PlaceOrderDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseObject"
                        }
                    }
                }
            }
        },
        "/orders/{id}": {
            "get": {
                "description": "Get Single Order by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Get Single Order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/orders/{id}/cancel": {
            "put": {
                "description": "Cancel Order with a given Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Cancel Order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/orders/{id}/status": {
            "put": {
                "description": "Update Order Status with a given Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Update Order Status",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "data to update order status",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateOrderStatusDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/products": {
            "get": {
                "description": "Gets All products",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get All Products",
                "parameters": [
                    {
                        "description": "data to query for all ",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.APIPagingDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Create new Product",
                "parameters": [
                    {
                        "description": "data to create new product",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateProductDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseObject"
                        }
                    }
                }
            }
        },
        "/products/{id}": {
            "get": {
                "description": "Get Single product by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get Single product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "put": {
                "description": "Updates Product with a given Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Update Product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "data to update product with",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateProductDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete product by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Delete Product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.APIPagingDto": {
            "type": "object",
            "properties": {
                "direction": {
                    "type": "string"
                },
                "filter": {
                    "type": "string"
                },
                "limit": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "select": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "sort": {
                    "type": "string"
                }
            }
        },
        "models.CreateProductDto": {
            "type": "object",
            "required": [
                "currency",
                "description",
                "name",
                "price",
                "quantity"
            ],
            "properties": {
                "currency": {
                    "$ref": "#/definitions/models.Currency"
                },
                "description": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 4
                },
                "discount": {
                    "type": "integer"
                },
                "name": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 4
                },
                "price": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "models.Currency": {
            "type": "string",
            "enum": [
                "NGN"
            ],
            "x-enum-varnames": [
                "CURRENCY_NGN"
            ]
        },
        "models.OrderStatus": {
            "type": "string",
            "enum": [
                "pending",
                "processing",
                "delivered",
                "cancelled",
                "shipped"
            ],
            "x-enum-varnames": [
                "PENDING",
                "PROCESSING",
                "DELIVERED",
                "CANCELLED",
                "SHIPPED"
            ]
        },
        "models.PlaceOrder": {
            "type": "object",
            "required": [
                "product_id",
                "quantity"
            ],
            "properties": {
                "product_id": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "models.PlaceOrderDto": {
            "type": "object",
            "required": [
                "currency"
            ],
            "properties": {
                "currency": {
                    "$ref": "#/definitions/models.Currency"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.PlaceOrder"
                    }
                }
            }
        },
        "models.ProductStatus": {
            "type": "string",
            "enum": [
                "in-stock",
                "not-in-stock",
                "sold-out"
            ],
            "x-enum-varnames": [
                "IN_STOCK",
                "NOT_IN_STOCK",
                "SOLD_OUT"
            ]
        },
        "models.ResponseObject": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "models.SignInDto": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.SignUpDto": {
            "type": "object",
            "required": [
                "email",
                "firstName",
                "lastName",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string",
                    "maxLength": 25,
                    "minLength": 2
                },
                "lastName": {
                    "type": "string",
                    "maxLength": 25,
                    "minLength": 2
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/models.UserRole"
                }
            }
        },
        "models.UpdateOrderStatusDto": {
            "type": "object",
            "required": [
                "status"
            ],
            "properties": {
                "status": {
                    "$ref": "#/definitions/models.OrderStatus"
                }
            }
        },
        "models.UpdateProductDto": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 4
                },
                "discount": {
                    "type": "integer"
                },
                "name": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 4
                },
                "price": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                },
                "status": {
                    "$ref": "#/definitions/models.ProductStatus"
                }
            }
        },
        "models.UserRole": {
            "type": "string",
            "enum": [
                "user",
                "admin"
            ],
            "x-enum-varnames": [
                "USER_ROLE_USER",
                "USER_ROLE_ADMIN"
            ]
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
