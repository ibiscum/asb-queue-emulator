// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "text/plain",
    "application/json"
  ],
  "produces": [
    "text/plain",
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "azure service bus emulator",
    "title": "azure service bus",
    "version": "1.0.0"
  },
  "paths": {
    "/{queueName}/messages": {
      "post": {
        "operationId": "sendMessage",
        "parameters": [
          {
            "type": "string",
            "description": "broker properties",
            "name": "brokerProperties",
            "in": "header"
          },
          {
            "type": "string",
            "description": "the queue name",
            "name": "queueName",
            "in": "path",
            "required": true
          },
          {
            "description": "message content",
            "name": "messageContent",
            "in": "body",
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Message successfully sent to queue or topic."
          },
          "400": {
            "description": "Bad request."
          },
          "401": {
            "description": "Authorization failure."
          },
          "403": {
            "description": "Quota exceeded or message too large."
          },
          "410": {
            "description": "Specified queue or topic does not exist."
          },
          "500": {
            "description": "Internal error."
          }
        }
      }
    },
    "/{queueName}/messages/head": {
      "post": {
        "operationId": "peekMessage",
        "parameters": [
          {
            "type": "string",
            "description": "the queue name",
            "name": "queueName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "Message successfully sent to queue or topic.",
            "schema": {
              "type": "string"
            },
            "headers": {
              "BrokerProperties": {
                "type": "string"
              },
              "Location": {
                "type": "string"
              }
            }
          },
          "204": {
            "description": "No messages available within the specified timeout period.."
          },
          "400": {
            "description": "Bad request."
          },
          "401": {
            "description": "Authorization failure."
          },
          "410": {
            "description": "Specified queue or subscription does not exist.."
          },
          "500": {
            "description": "Internal error."
          }
        }
      },
      "delete": {
        "operationId": "destructiveRead",
        "parameters": [
          {
            "type": "string",
            "description": "the queue name",
            "name": "queueName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Message successfully retrieved and deleted",
            "schema": {
              "type": "string"
            },
            "headers": {
              "BrokerProperties": {
                "type": "string"
              },
              "Location": {
                "type": "string"
              }
            }
          },
          "204": {
            "description": "No messages available within the specified timeout period.."
          },
          "400": {
            "description": "Bad request."
          },
          "401": {
            "description": "Authorization failure."
          },
          "410": {
            "description": "Specified queue or subscription does not exist.."
          },
          "500": {
            "description": "Internal error."
          }
        }
      }
    },
    "/{queueName}/messages/{messageId}/{lockToken}": {
      "delete": {
        "operationId": "deleteMessage",
        "parameters": [
          {
            "type": "string",
            "description": "the queue name",
            "name": "queueName",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "The ID of the message to be deleted as returned in BrokerProperties{MessageId} by the Peek Message operation.",
            "name": "messageId",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "The token of the lock of the message to be deleted as returned by the Peek Message operation in BrokerProperties{LockToken}.",
            "name": "lockToken",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Message successfully deleted..",
            "headers": {
              "BrokerProperties": {
                "type": "string"
              },
              "Location": {
                "type": "string"
              }
            }
          },
          "204": {
            "description": "No messages available within the specified timeout period.."
          },
          "400": {
            "description": "Bad request."
          },
          "401": {
            "description": "Authorization failure."
          },
          "404": {
            "description": "No message was found with the specified MessageId or LockToken."
          },
          "410": {
            "description": "Specified queue or subscription does not exist."
          },
          "500": {
            "description": "Internal error."
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json",
    "text/plain"
  ],
  "produces": [
    "application/json",
    "text/plain"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "azure service bus emulator",
    "title": "azure service bus",
    "version": "1.0.0"
  },
  "paths": {
    "/{queueName}/messages": {
      "post": {
        "operationId": "sendMessage",
        "parameters": [
          {
            "type": "string",
            "description": "broker properties",
            "name": "brokerProperties",
            "in": "header"
          },
          {
            "type": "string",
            "description": "the queue name",
            "name": "queueName",
            "in": "path",
            "required": true
          },
          {
            "description": "message content",
            "name": "messageContent",
            "in": "body",
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Message successfully sent to queue or topic."
          },
          "400": {
            "description": "Bad request."
          },
          "401": {
            "description": "Authorization failure."
          },
          "403": {
            "description": "Quota exceeded or message too large."
          },
          "410": {
            "description": "Specified queue or topic does not exist."
          },
          "500": {
            "description": "Internal error."
          }
        }
      }
    },
    "/{queueName}/messages/head": {
      "post": {
        "operationId": "peekMessage",
        "parameters": [
          {
            "type": "string",
            "description": "the queue name",
            "name": "queueName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "Message successfully sent to queue or topic.",
            "schema": {
              "type": "string"
            },
            "headers": {
              "BrokerProperties": {
                "type": "string"
              },
              "Location": {
                "type": "string"
              }
            }
          },
          "204": {
            "description": "No messages available within the specified timeout period.."
          },
          "400": {
            "description": "Bad request."
          },
          "401": {
            "description": "Authorization failure."
          },
          "410": {
            "description": "Specified queue or subscription does not exist.."
          },
          "500": {
            "description": "Internal error."
          }
        }
      },
      "delete": {
        "operationId": "destructiveRead",
        "parameters": [
          {
            "type": "string",
            "description": "the queue name",
            "name": "queueName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Message successfully retrieved and deleted",
            "schema": {
              "type": "string"
            },
            "headers": {
              "BrokerProperties": {
                "type": "string"
              },
              "Location": {
                "type": "string"
              }
            }
          },
          "204": {
            "description": "No messages available within the specified timeout period.."
          },
          "400": {
            "description": "Bad request."
          },
          "401": {
            "description": "Authorization failure."
          },
          "410": {
            "description": "Specified queue or subscription does not exist.."
          },
          "500": {
            "description": "Internal error."
          }
        }
      }
    },
    "/{queueName}/messages/{messageId}/{lockToken}": {
      "delete": {
        "operationId": "deleteMessage",
        "parameters": [
          {
            "type": "string",
            "description": "the queue name",
            "name": "queueName",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "The ID of the message to be deleted as returned in BrokerProperties{MessageId} by the Peek Message operation.",
            "name": "messageId",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "The token of the lock of the message to be deleted as returned by the Peek Message operation in BrokerProperties{LockToken}.",
            "name": "lockToken",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Message successfully deleted..",
            "headers": {
              "BrokerProperties": {
                "type": "string"
              },
              "Location": {
                "type": "string"
              }
            }
          },
          "204": {
            "description": "No messages available within the specified timeout period.."
          },
          "400": {
            "description": "Bad request."
          },
          "401": {
            "description": "Authorization failure."
          },
          "404": {
            "description": "No message was found with the specified MessageId or LockToken."
          },
          "410": {
            "description": "Specified queue or subscription does not exist."
          },
          "500": {
            "description": "Internal error."
          }
        }
      }
    }
  }
}`))
}
