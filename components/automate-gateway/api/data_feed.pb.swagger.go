package api

func init() {
	Swagger.Add("data_feed", `{
  "swagger": "2.0",
  "info": {
    "title": "api/external/data_feed/data_feed.proto",
    "version": "version not set"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/api/v0/datafeed/destination": {
      "post": {
        "summary": "Add a datafeed destination",
        "description": "Adds a datafeed destination. Requires values for name, url, and secret.\nThe secret is the id returned from creating a secret using the secrets api.\n\nExample:\n` + "`" + `` + "`" + `` + "`" + `\n{\n\"name\": \"my datafeed destination\",\n\"url\": \"https://my.server.com/dataingest\",\n\"secret\":\"42b369f1-9d3f-44b3-bcf8-a9a09d6bd4bb\"\n}\n` + "`" + `` + "`" + `` + "`" + `\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\ndatafeed:destination:add\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "AddDestination",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.datafeed.AddDestinationResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.datafeed.AddDestinationRequest"
            }
          }
        ],
        "tags": [
          "DatafeedService"
        ]
      }
    },
    "/api/v0/datafeed/destination/{id}": {
      "get": {
        "summary": "Get a datafeed destination",
        "description": "Gets a datafeed destination the ID of the destination.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\ndatafeed:destination:get\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "GetDestination",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.datafeed.GetDestinationResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "int64"
          }
        ],
        "tags": [
          "DatafeedService"
        ]
      },
      "delete": {
        "summary": "Delete a datafeed destination",
        "description": "Deletes a datafeed destination given the ID of the destination.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\ndestination:destination:delete\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "DeleteDestination",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.datafeed.DeleteDestinationResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "int64"
          }
        ],
        "tags": [
          "DatafeedService"
        ]
      },
      "patch": {
        "summary": "Update a datafeed destination",
        "description": "Updates a datafeed destination.\nThis is a PATCH operation, meaning the details sent in will override/replace those stored in the DB.\nThis will update the name, url or secret for the destination\n\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\ndestination:destination:update\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "UpdateDestination",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.datafeed.UpdateDestinationResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.datafeed.UpdateDestinationRequest"
            }
          }
        ],
        "tags": [
          "DatafeedService"
        ]
      }
    },
    "/api/v0/datafeed/destinations": {
      "post": {
        "summary": "List Destinations",
        "description": "Returns a list of all datafeed destinations.\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\ndatafeed:destinations:list\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "ListDestinations",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.datafeed.ListDestinationResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.datafeed.ListDestinationRequest"
            }
          }
        ],
        "tags": [
          "DatafeedService"
        ]
      }
    },
    "/api/v0/datafeed/destinations/test": {
      "post": {
        "summary": "Test a datafeed destination",
        "description": "Tests a datafeed destination endpoint with the details provided. Requires values for name, url, and secret.\nThe secret is the id returned from creating a secret using the secrets api.\nAccepts either url, secret or url, username_password\n\nExamples:\n` + "`" + `` + "`" + `` + "`" + `\n{\n\"url\": \"https://my.server.com/dataingest\",\n\"secret\":\"42b369f1-9d3f-44b3-bcf8-a9a09d6bd4bb\"\n}\n\n{\n\"url\": \"https://my.server.com/dataingest\",\n\"username_password\": {\n\"username\": \"muyuser\",\n\"password\": \"mypassword\"\n}\n}\n` + "`" + `` + "`" + `` + "`" + `\n\nAuthorization Action:\n` + "`" + `` + "`" + `` + "`" + `\ndatafeed:destinations:test\n` + "`" + `` + "`" + `` + "`" + `",
        "operationId": "TestDestination",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.datafeed.TestDestinationResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.datafeed.URLValidationRequest"
            }
          }
        ],
        "tags": [
          "DatafeedService"
        ]
      }
    }
  },
  "definitions": {
    "chef.automate.api.datafeed.AddDestinationRequest": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "format": "int64"
        },
        "name": {
          "type": "string"
        },
        "url": {
          "type": "string"
        },
        "secret": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.datafeed.AddDestinationResponse": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "format": "int64"
        },
        "name": {
          "type": "string"
        },
        "url": {
          "type": "string"
        },
        "secret": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.datafeed.DeleteDestinationResponse": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "format": "int64"
        },
        "name": {
          "type": "string"
        },
        "url": {
          "type": "string"
        },
        "secret": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.datafeed.GetDestinationResponse": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "format": "int64"
        },
        "name": {
          "type": "string"
        },
        "url": {
          "type": "string"
        },
        "secret": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.datafeed.ListDestinationRequest": {
      "type": "object"
    },
    "chef.automate.api.datafeed.ListDestinationResponse": {
      "type": "object",
      "properties": {
        "destinations": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.datafeed.GetDestinationResponse"
          }
        }
      }
    },
    "chef.automate.api.datafeed.SecretId": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.datafeed.TestDestinationResponse": {
      "type": "object",
      "properties": {
        "success": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "chef.automate.api.datafeed.URLValidationRequest": {
      "type": "object",
      "properties": {
        "url": {
          "type": "string"
        },
        "username_password": {
          "$ref": "#/definitions/chef.automate.api.datafeed.UsernamePassword"
        },
        "secret_id": {
          "$ref": "#/definitions/chef.automate.api.datafeed.SecretId"
        }
      }
    },
    "chef.automate.api.datafeed.UpdateDestinationRequest": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "url": {
          "type": "string"
        },
        "secret": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.datafeed.UpdateDestinationResponse": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "format": "int64"
        },
        "name": {
          "type": "string"
        },
        "url": {
          "type": "string"
        },
        "secret": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.datafeed.UsernamePassword": {
      "type": "object",
      "properties": {
        "username": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    }
  }
}
`)
}
