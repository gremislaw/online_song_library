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
        "/api/v1/create": {
            "post": {
                "description": "Creates a new song with the provided details.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a new song.",
                "operationId": "create-song",
                "parameters": [
                    {
                        "description": "The song details to be created",
                        "name": "Song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Song"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "The song was created successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Invalid request or bad request to API",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/delete/{id}": {
            "delete": {
                "description": "Deletes a song by its ID.",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete a song by its ID.",
                "operationId": "delete-song",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The ID of the song to be deleted",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "The song was deleted successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Invalid ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Song not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/get": {
            "get": {
                "description": "Retrieves a list of songs with optional filtering and pagination.",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a list of songs",
                "operationId": "get-songs",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Filter by group name",
                        "name": "group",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by song name",
                        "name": "song",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Limit the number of results",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "Offset for pagination",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "A map containing the list of songs",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/get-info": {
            "get": {
                "description": "Retrieves additional information about a song from an external API based on the group and song title.",
                "produces": [
                    "application/json"
                ],
                "summary": "Get additional information about a song from an external API.",
                "operationId": "get-info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The name of the group.",
                        "name": "group",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The title of the song.",
                        "name": "song",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "The detailed information about the song.",
                        "schema": {
                            "$ref": "#/definitions/models.SongDetail"
                        }
                    },
                    "400": {
                        "description": "Invalid request or bad request to API.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/get/{id}": {
            "get": {
                "description": "Retrieves the text of a song by its ID with optional pagination.",
                "produces": [
                    "application/json"
                ],
                "summary": "Get the text of a song by its ID.",
                "operationId": "get-song-text",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The ID of the song",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Limit the number of lines",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "Offset for pagination",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "The text of the song",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Invalid ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Song not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/update/{id}": {
            "patch": {
                "description": "Updates an existing song with the provided details.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update an existing song by its ID.",
                "operationId": "update-song",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The ID of the song to be updated",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "The updated song details",
                        "name": "Song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Song"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "The song was updated successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Invalid request or bad request to API",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Song not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Song": {
            "description": "Song model",
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "Created date of the song",
                    "type": "string"
                },
                "group": {
                    "description": "Group name",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the song",
                    "type": "integer"
                },
                "link": {
                    "description": "Link to the song",
                    "type": "string"
                },
                "releaseDate": {
                    "description": "Release date of the song",
                    "type": "string"
                },
                "song": {
                    "description": "Song title",
                    "type": "string"
                },
                "text": {
                    "description": "Song text",
                    "type": "string"
                },
                "updatedAt": {
                    "description": "Updated date of the song",
                    "type": "string"
                }
            }
        },
        "models.SongDetail": {
            "description": "SongDetail model",
            "type": "object",
            "properties": {
                "link": {
                    "description": "Text of the song",
                    "type": "string"
                },
                "releaseDate": {
                    "description": "Link to the song",
                    "type": "string"
                },
                "text": {
                    "description": "Release date of the song",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Online Song Library API",
	Description:      "Swagger API for Golang Online Song Library",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
