package docs

import (
	"github.com/swaggo/swag"
)

type s struct{}

func (s *s) ReadDoc() string {
	return docs
}
func init() {
	swag.Register(swag.Name, &s{})
}

const docs = `
{
  "swagger": "2.0",
  "info": {
    "description": "API specification for APD Store application",
    "title": "APD Store",
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0"
  },
  "schemes": [
    "http"
  ],
  "host": "137.116.136.106",
  "basePath": "/api",
  "tags": [
    {
      "name": "product",
      "description": "Product that sold by merchant and can be bought by customers"
    },
    {
      "name": "user",
      "description": "User in this application"
    },
    {
      "name": "merchant",
      "description": "Merchant that sold item to users"
    },
    {
      "name": "transaction",
      "description": "Transaction that indicate customer buying status"
    },
    {
      "name": "admin",
      "description": "Admin that decide the transaction and proposal should be accepted or not"
    },
    {
      "name": "transfer",
      "description": "Transfer that done by merchant from or to their bank account"
    },
    {
      "name": "auth",
      "description": "Handle authentication and authorization"
    }
  ],
  "security": [
    JWT:[]
  ],
  "securityDefinitions": {
    "JWT": {
      "description": "",
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  },
  "paths": {
    "/product": {
      "post": {
        "tags": [
          "product"
        ],
        "description": "Creating product that can be sold for user",
        "consumes": [
          "multipart/form-data"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Creating product for merchant",
        "parameters": [
          {
            "type": "string",
            "description": "Name for product",
            "name": "name",
            "in": "formData",
            "required": true
          },
          {
            "type": "file",
            "description": "Image for product",
            "name": "image",
            "in": "formData",
            "required": false
          },
          {
            "type": "integer",
            "description": "Price for product",
            "name": "price",
            "in": "formData",
            "required": true
          },
          {
            "type": "string",
            "description": "Description for product",
            "name": "description",
            "in": "formData",
            "required": false
          },
          {
            "type": "integer",
            "description": "Stock that available for this product",
            "name": "stock",
            "in": "formData",
            "required": true
          },
          {
            "type": "string",
            "description": "Merchant id that selling this product",
            "name": "merchantId",
            "in": "formData",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/createdResponse"
            }
          },
          "404": {
            "description": "Id is invalid",
            "schema": {
              "$ref": "#/definitions/notFoundResponse"
            }
          }
        }
      },
      "get": {
        "tags": [
          "product"
        ],
        "summary": "Get product by id",
        "produces": [
          "application/json"
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/getProductByIdResponse"
            }
          },
          "404": {
            "description": "Id is invalid",
            "schema": {
              "$ref": "#/definitions/notFoundResponse"
            }
          }
        }
      },
      "put": {
        "tags": [
          "product"
        ],
        "summary": "Update product data",
        "consumes": [
          "multipart/form-data"
        ],
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "type": "string",
            "description": "Id for product",
            "name": "product_id",
            "in": "formData",
            "required": true
          },
          {
            "type": "string",
            "description": "Name for product",
            "name": "name",
            "in": "formData",
            "required": true
          },
          {
            "type": "file",
            "description": "Image for product",
            "name": "image",
            "in": "formData",
            "required": false
          },
          {
            "type": "integer",
            "description": "Price for product",
            "name": "price",
            "in": "formData",
            "required": true
          },
          {
            "type": "string",
            "description": "Description for product",
            "name": "description",
            "in": "formData",
            "required": false
          },
          {
            "type": "integer",
            "description": "Stock that available for this product",
            "name": "stock",
            "in": "formData",
            "required": true
          },
          {
            "type": "string",
            "description": "Merchant id that selling this product",
            "name": "merchantId",
            "in": "formData",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/successResponse"
            }
          },
          "404": {
            "description": "Id is invalid",
            "schema": {
              "$ref": "#/definitions/notFoundResponse"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "product"
        ],
        "summary": "Delete product by id",
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "in": "query",
            "type": "string",
            "name": "productId",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/successResponse"
            }
          },
          "404": {
            "description": "Id is invalid",
            "schema": {
              "$ref": "#/definitions/notFoundResponse"
            }
          }
        }
      }
    },
    "/products": {
      "get": {
        "tags": [
          "product"
        ],
        "summary": "Get all products",
        "produces": [
          "application/json"
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/getProductsResponse"
            }
          },
          "404": {
            "description": "Id is invalid",
            "schema": {
              "$ref": "#/definitions/notFoundResponse"
            }
          }
        }
      }
    },
    "/merchant/products": {
      "get": {
        "tags": [
          "product"
        ],
        "parameters": [
          {
            "in": "query",
            "type": "string",
            "name": "merchantId",
            "required": true
          }
        ],
        "summary": "Get all products that sold by merchant",
        "produces": [
          "application/json"
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/getProductsResponse"
            }
          },
          "404": {
            "description": "Id is invalid",
            "schema": {
              "$ref": "#/definitions/notFoundResponse"
            }
          }
        }
      }
    },
    "/transaction": {
      "post": {
        "tags": [
          "transaction"
        ],
        "summary": "Create a transaction",
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "request",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/createTransactionRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/successResponse"
            }
          },
          "404": {
            "description": "Id is invalid",
            "schema": {
              "$ref": "#/definitions/notFoundResponse"
            }
          }
        }
      },
      "get": {
        "tags": [
          "transaction"
        ],
        "summary": "Get transaction by id",
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "request",
            "type": "string",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/getTransactionByIdResponse"
            }
          },
          "404": {
            "description": "Id is invalid",
            "schema": {
              "$ref": "#/definitions/notFoundResponse"
            }
          }
        }
      }
    },
    "/transaction/status": {
      "post": {
        "tags": [
          "transaction"
        ],
        "summary": "Update transaction status",
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "request",
            "description": "For status the values are: 0 on cart, 1 waiting payment, 2 waiting confirmation, 3 declined, 4 waiting delivery, 5 accepted",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/updateTransactionRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/successResponse"
            }
          },
          "404": {
            "description": "Id is invalid",
            "schema": {
              "$ref": "#/definitions/notFoundResponse"
            }
          }
        }
      }
    },
    "/transactions/history": {
      "get": {
        "tags": [
          "transaction"
        ],
        "summary": "Get transaction history by user id",
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "user_id",
            "in": "query",
            "type": "string"
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/getTransactionHistoryResponse"
            }
          },
          "404": {
            "description": "Id is invalid",
            "schema": {
              "$ref": "#/definitions/notFoundResponse"
            }
          }
        }
      }
    },
    "/transactions/request": {
      "get": {
        "tags": [
          "transaction"
        ],
        "summary": "Get merchant request item list",
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "merchantId",
            "in": "query",
            "type": "string"
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/getTransactionHistoryResponse"
            }
          },
          "404": {
            "description": "Id is invalid",
            "schema": {
              "$ref": "#/definitions/notFoundResponse"
            }
          }
        }
      }
    },
    "/transaction/payment": {
      "post": {
        "tags": [
          "transaction"
        ],
        "summary": "Pay for the transaction",
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "request",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/paymentRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/successResponse"
            }
          },
          "404": {
            "description": "Id is invalid",
            "schema": {
              "$ref": "#/definitions/notFoundResponse"
            }
          }
        }
      }
    },
    "/transfer": {
      "post": {
        "tags": [
          "transfer"
        ],
        "summary": "Transfer in or out to bank",
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "request",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/transferRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/successResponse"
            }
          },
          "404": {
            "description": "Id is invalid",
            "schema": {
              "$ref": "#/definitions/notFoundResponse"
            }
          }
        }
      }
    },
    "/transfers": {
      "get": {
        "tags": [
          "transfer"
        ],
        "summary": "Transfer in or out from bank",
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "merchantId",
            "in": "query",
            "type": "string"
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/getTransferResponse"
            }
          },
          "404": {
            "description": "Id is invalid",
            "schema": {
              "$ref": "#/definitions/notFoundResponse"
            }
          }
        }
      }
    },
    "/merchant": {
      "post": {
        "tags": [
          "merchant"
        ],
        "summary": "Create a merchant",
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "request",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/merchantRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/successResponse"
            }
          },
          "404": {
            "description": "Id is invalid",
            "schema": {
              "$ref": "#/definitions/notFoundResponse"
            }
          }
        }
      },
      "get": {
        "tags": [
          "merchant"
        ],
        "summary": "Get Merchant by ID",
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "merchantId",
            "in": "query",
            "type": "string"
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/getMerchantByIdResponse"
            }
          },
          "404": {
            "description": "Id is invalid",
            "schema": {
              "$ref": "#/definitions/notFoundResponse"
            }
          }
        }
      },
      "put": {
        "tags": [
          "merchant"
        ],
        "summary": "Update a merchant",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "request",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/updateMerchantRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/successResponse"
            }
          },
          "404": {
            "description": "Id is invalid",
            "schema": {
              "$ref": "#/definitions/notFoundResponse"
            }
          }
        }
      }
    },
    "/merchants": {
      "get": {
        "tags": [
          "merchant"
        ],
        "summary": "Get list of all merchants",
        "produces": [
          "application/json"
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/getMerchantsResponse"
            }
          },
          "404": {
            "description": "Id is invalid",
            "schema": {
              "$ref": "#/definitions/notFoundResponse"
            }
          }
        }
      }
    },
    "/merchant/balance": {
      "get": {
        "tags": [
          "merchant"
        ],
        "summary": "Get Merchant's Balance by ID",
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "merchantId",
            "in": "query",
            "type": "string"
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/getMerchantBalanceResponse"
            }
          },
          "404": {
            "description": "Id is invalid",
            "schema": {
              "$ref": "#/definitions/notFoundResponse"
            }
          }
        }
      }
    },
    "/merchant/status": {
      "put": {
        "tags": [
          "merchant"
        ],
        "summary": "Update Merchant's Approval Status",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "request",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/updateMerchantApprovalStatusRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/successResponse"
            }
          },
          "404": {
            "description": "Id is invalid",
            "schema": {
              "$ref": "#/definitions/notFoundResponse"
            }
          }
        }
      }
    },
    "/auth/register": {
      "post": {
        "tags": [
          "auth"
        ],
        "description": "Create User",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "request",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/userRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/createdResponse"
            }
          },
          "404": {
            "description": "Id is invalid",
            "schema": {
              "$ref": "#/definitions/notFoundResponse"
            }
          }
        }
      }
    },
    "/user": {
      "put": {
        "tags": [
          "user"
        ],
        "description": "Update user",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "request",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/updateUserRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/createdResponse"
            }
          },
          "404": {
            "description": "Id is invalid",
            "schema": {
              "$ref": "#/definitions/notFoundResponse"
            }
          }
        }
      }
    },
    "/auth/login": {
      "post": {
        "tags": [
          "auth"
        ],
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "request",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/loginRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/successResponse"
            }
          },
          "401": {
            "description": "unauthenticated",
            "schema": {
              "$ref": "#/definitions/unauthenticatedResponse"
            }
          }
        }
      }
    },
    "/user/admin": {
      "post": {
        "tags": [
          "admin"
        ],
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "request",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/registerAdminRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/successResponse"
            }
          },
          "400": {
            "description": "bad request",
            "schema": {
              "$ref": "#/definitions/badRequestResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "createTransactionRequest": {
      "type": "object",
      "properties": {
        "bank_number": {
          "type": "string",
          "example": "123"
        },
        "bank_name": {
          "type": "string",
          "example": "Bank World"
        },
        "amount": {
          "type": "integer",
          "example": 10000
        },
        "user_id": {
          "type": "string",
          "example": "3efa0889-c74b-4498-9134-5733bfdbd8a7"
        }
      }
    },
    "updateTransactionRequest": {
      "type": "object",
      "properties": {
        "transaction_id": {
          "description": "asdasd",
          "type": "string",
          "example": "3efa0889-c74b-4498-9134-5733bfdbd8a7"
        },
        "status": {
          "type": "integer",
          "example": 1,
          "enum": [
            0,
            1,
            2,
            3,
            4,
            5
          ]
        }
      }
    },
    "merchantRequest": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "example": "SEA Store"
        },
        "user_id": {
          "type": "string",
          "example": "3efa0889-c74b-4498-9134-5733bfdbd8a7"
        },
        "brand": {
          "type": "string",
          "example": "Brand"
        },
        "address": {
          "type": "string",
          "example": "Downing Street"
        }
      }
    },
    "paymentRequest": {
      "type": "object",
      "properties": {
        "customer_id": {
          "type": "string",
          "example": "3efa0889-c74b-4498-9134-5733bfdbd8a7"
        },
        "bank_number": {
          "type": "string",
          "example": "123"
        },
        "bank_name": {
          "type": "string",
          "example": "Bank name"
        },
        "transaction_id": {
          "type": "string",
          "example": "3efa0889-c74b-4498-9134-5733bfdbd8a7"
        }
      }
    },
    "transferRequest": {
      "type": "object",
      "properties": {
        "merchant_id": {
          "type": "string",
          "example": "3efa0889-c74b-4498-9134-5733bfdbd8a7"
        },
        "bank_number": {
          "type": "string",
          "example": "123"
        },
        "bank_name": {
          "type": "string",
          "example": "Bank name"
        },
        "amount": {
          "type": "integer",
          "example": -100000
        }
      }
    },
    "userRequest": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string",
          "example": "john@example.xyz"
        },
        "password": {
          "type": "string",
          "example": "********"
        },
        "name": {
          "type": "string",
          "example": "John Doe"
        }
      }
    },
    "createdResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "example": 201
        },
        "message": {
          "type": "string",
          "example": "success"
        }
      }
    },
    "notFoundResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "example": 404
        },
        "message": {
          "type": "string",
          "example": "not found"
        }
      }
    },
    "successResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "example": 200
        },
        "message": {
          "type": "string",
          "example": "success"
        }
      }
    },
    "badRequestResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "example": 400
        },
        "message": {
          "type": "string",
          "example": "bad request"
        }
      }
    },
    "unauthenticatedResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "example": 401
        },
        "message": {
          "type": "string",
          "example": "unauthenticated"
        }
      }
    },
    "getProductsResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "example": 200
        },
        "message": {
          "type": "string",
          "example": "success"
        },
        "data": {
          "type": "object",
          "properties": {
            "products": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/model.product"
              }
            }
          }
        }
      }
    },
    "getProductByIdResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "example": 200
        },
        "message": {
          "type": "string",
          "example": "success"
        },
        "data": {
          "$ref": "#/definitions/model.product"
        }
      }
    },
    "getTransactionByIdResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "example": 200
        },
        "message": {
          "type": "string",
          "example": "success"
        },
        "data": {
          "$ref": "#/definitions/model.transaction"
        }
      }
    },
    "getTransactionHistoryResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "example": 200
        },
        "message": {
          "type": "string",
          "example": "success"
        },
        "data": {
          "type": "object",
          "properties": {
            "transactions": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/model.transaction"
              }
            }
          }
        }
      }
    },
    "getTransferResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "example": 200
        },
        "message": {
          "type": "string",
          "example": "success"
        },
        "data": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/model.transfer"
          }
        }
      }
    },
    "getMerchantsResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "example": 200
        },
        "message": {
          "type": "string",
          "example": "success"
        },
        "data": {
          "type": "object",
          "properties": {
            "products": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/model.merchant"
              }
            }
          }
        }
      }
    },
    "getMerchantByIdResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "example": 200
        },
        "message": {
          "type": "string",
          "example": "success"
        },
        "data": {
          "$ref": "#/definitions/model.merchant"
        }
      }
    },
    "getMerchantBalanceResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "example": 200
        },
        "message": {
          "type": "string",
          "example": "success"
        },
        "data": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "balance": {
                "type": "integer"
              }
            }
          }
        }
      }
    },
    "model.product": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "id for product in uuid format",
          "example": "867ce2ce-77ec-44da-9404-d328db73e729"
        },
        "name": {
          "type": "string",
          "description": "Name for product",
          "example": "T shirt"
        },
        "image": {
          "type": "string",
          "description": "File path for product image",
          "example": "http://google.com"
        },
        "price": {
          "type": "integer",
          "description": "Price for product",
          "example": "10000"
        },
        "description": {
          "type": "string",
          "description": "Description for product",
          "example": "This is a comfortable shirt"
        },
        "stock": {
          "type": "integer",
          "description": "Stock that available for this product",
          "example": "3"
        },
        "merchant_id": {
          "type": "string",
          "description": "Merchant id that selling this product",
          "example": "2538dae4-3bed-47f2-898c-e0b80bb2816e"
        },
        "created_at": {
          "type": "string",
          "format": "date-time",
          "description": "Timestamp that indicate when this product is created",
          "example": "2020-09-08T17:48:56+0000"
        },
        "updated_at": {
          "type": "string",
          "format": "date-time",
          "description": "Timestamp that indicate when this product is updated",
          "example": "2020-09-08T17:48:56+0000"
        },
        "deleted_at": {
          "type": "string",
          "format": "date-time",
          "description": "Timestamp that indicate when this product is deleted",
          "example": "2020-09-08T17:48:56+0000"
        }
      }
    },
    "model.transaction": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "id for transaction in uuid format",
          "example": "867ce2ce-77ec-44da-9404-d328db73e729"
        },
        "status": {
          "type": "string",
          "enum": [
            "accepted",
            "onprogress",
            "declined"
          ],
          "description": "status for the trasaction",
          "example": "accepted"
        },
        "bank_number": {
          "type": "string",
          "description": "Bank number that used for transaction",
          "example": "123456789"
        },
        "bank_name": {
          "type": "string",
          "description": "Bank name that used for transaction",
          "example": "Bank World"
        },
        "amount": {
          "type": "integer",
          "description": "Amount of money for this transaction",
          "example": 20000
        },
        "user_id": {
          "type": "string",
          "description": "Buyer that doing the transaction",
          "example": "867ce2ce-77ec-44da-9404-d328db73e729"
        }
      }
    },
    "model.transfer": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "id for transfer in uuid format",
          "example": "867ce2ce-77ec-44da-9404-d328db73e729"
        },
        "bank_number": {
          "type": "string",
          "description": "Bank number that used for transfer",
          "example": "123456789"
        },
        "bank_name": {
          "type": "string",
          "description": "Bank name that used for transfer",
          "example": "Bank World"
        },
        "amount": {
          "type": "integer",
          "description": "Amount of money for this transfer",
          "example": 20000
        },
        "merchant_id": {
          "type": "string",
          "description": "Indicate merchant that do this transfer",
          "example": "867ce2ce-77ec-44da-9404-d328db73e729"
        }
      }
    },
    "model.merchant": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "id for merchant in uuid format",
          "example": "867ce2ce-77ec-44da-9404-d328db73e729"
        },
        "name": {
          "type": "string",
          "description": "Merchant's name",
          "example": "SEA store"
        },
        "balance": {
          "type": "integer",
          "description": "merchant's balance",
          "example": 500000
        },
        "userId": {
          "type": "string",
          "description": "Indicate user that having this merchant",
          "example": "867ce2ce-77ec-44da-9404-d328db73e729"
        },
        "brand": {
          "type": "string",
          "description": "the brand",
          "example": "Brand"
        },
        "address": {
          "type": "string",
          "description": "merchant's address",
          "example": "Downing Street"
        },
        "approval": {
          "type": "string",
          "enum": [
            "waiting",
            "accepted",
            "others",
            "declined"
          ],
          "description": "status for the merchant",
          "example": "accepted"
        }
      }
    },
    "model.user": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "id for user in uuid format",
          "example": "867ce2ce-77ec-44da-9404-d328db73e729"
        },
        "name": {
          "type": "string",
          "description": "User's name",
          "example": "John Doe"
        },
        "email": {
          "type": "string",
          "description": "User's email",
          "example": "john@example.xyz"
        },
        "password": {
          "type": "string",
          "description": "User's password",
          "example": "********"
        },
        "phone": {
          "type": "string",
          "description": "user's phone",
          "example": "62181811781"
        },
        "role": {
          "type": "string",
          "enum": [
            "customer",
            "merchant",
            "admin",
            "other"
          ],
          "description": "User's role in the system"
        },
        "createdAt": {
          "type": "string",
          "description": "The timestamp indicates when the user created",
          "format": "date-time",
          "example": "2020-09-08T17:48:56+0000"
        },
        "updatedAt": {
          "type": "string",
          "description": "The timestamp indicates when the user updated",
          "format": "date-time",
          "example": "2020-09-08T17:48:56+0000"
        },
        "deletedAt": {
          "type": "string",
          "description": "The timestamp indicates when the user deleted",
          "format": "date-time",
          "example": "2020-09-08T17:48:56+0000"
        }
      }
    },
    "updateMerchantRequest": {
      "type": "object",
      "properties": {
        "user_id": {
          "type": "string",
          "example": "3efa0889-c74b-4498-9134-5733bfdbd8a7"
        },
        "name": {
          "type": "string",
          "example": "SEA Store"
        },
        "brand": {
          "type": "string",
          "example": "Brand"
        },
        "address": {
          "type": "string",
          "example": "Downing Street"
        }
      }
    },
    "updateMerchantApprovalStatusRequest": {
      "type": "object",
      "properties": {
        "approval": {
          "type": "string",
          "enum": [
            "waiting",
            "accepted",
            "others",
            "declined"
          ],
          "description": "status for the merchant",
          "example": "accepted"
        },
        "merchant_id": {
          "type": "string",
          "description": "Indicate user that having this merchant",
          "example": "867ce2ce-77ec-44da-9404-d328db73e729"
        }
      }
    },
    "updateUserRequest": {
      "type": "object",
      "properties": {
        "old_password": {
          "type": "string"
        },
        "new_password": {
          "type": "string"
        },
        "old_email": {
          "type": "string"
        },
        "new_email": {
          "type": "string"
        },
        "user_id": {
          "type": "string",
          "description": "id for user in uuid format",
          "example": "867ce2ce-77ec-44da-9404-d328db73e729"
        }
      }
    },
    "loginRequest": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "registerAdminRequest": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        },
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    }
  }
}
`
