package swagger

import (
	"encoding/json"
	"github.com/go-openapi/spec"
	"net/http"
)

type SwaggerServer struct{}

func New() *SwaggerServer {
	return &SwaggerServer{}
}

// todo: build definitions from: https://github.com/emicklei/go-restful-openapi/blob/master/definition_builder.go
// todo: build paths from: https://github.com/emicklei/go-restful-openapi/blob/master/build_path.go

func (s *SwaggerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	swagger := &spec.Swagger{
		SwaggerProps: spec.SwaggerProps{
			ID:       "id",
			Consumes: []string{"application/json"},
			Produces: []string{"application/json"},
			Schemes:  []string{"http"},
			Swagger:  "2.0",
			Info: &spec.Info{
				InfoProps: spec.InfoProps{
					Description:    "Description",
					Title:          "Title",
					TermsOfService: "Terms of Service",
					Contact: &spec.ContactInfo{
						Name:  "First Last",
						URL:   "contact-url.com",
						Email: "user@email.com",
					},
					License: &spec.License{
						Name: "MIT",
						URL:  "license-url.com",
					},
					Version: "api version",
				},
			},
			Host:     "host",
			BasePath: "basepath",
			Paths: &spec.Paths{
				Paths: map[string]spec.PathItem{
					"/users": {
						PathItemProps: spec.PathItemProps{
							Get: &spec.Operation{
								OperationProps: spec.OperationProps{
									Description:  "description",
									Consumes:     []string{},
									Produces:     []string{},
									Schemes:      []string{},
									Tags:         []string{},
									ExternalDocs: &spec.ExternalDocumentation{},
									Summary:      "",
									ID:           "",
									Deprecated:   false,
									Security:     []map[string][]string{},
									Parameters:   []spec.Parameter{},
									Responses:    &spec.Responses{},
								},
							},
						},
					},
				},
			},
			Definitions: spec.Definitions{
				"user": spec.Schema{
					
				},	
			},
		},
	}

	/*
		body, err := swagger.MarshalJSON()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	*/

	body, err := json.MarshalIndent(swagger, " ", "    ")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(body)
}

// GET Examples at: https://github.com/OAI/OpenAPI-Specification/tree/master/examples/v2.0/json

var swaggerJSON string = `
{
  "swaggerVersion": "1.2",
  "basePath": "http://localhost:8000/greetings",
  "apis": [
    {
      "path": "/hello/{subject}",
      "operations": [
        {
          "method": "GET",
          "summary": "Greet our subject with hello!",
          "type": "string",
          "nickname": "helloSubject",
          "parameters": [
            {
              "name": "subject",
              "description": "The subject to be greeted.",
              "required": true,
              "type": "string",
              "paramType": "path"
            }
          ]
        }
      ]
    }
  ],
  "models": {}
}
`

var swaggerJSON2 string = `
{
  "swagger": "2.0",
  "info": {
    "version": "1.0.0",
    "title": "Swagger Petstore",
    "description": "A sample API that uses a petstore as an example to demonstrate features in the swagger-2.0 specification",
    "termsOfService": "http://swagger.io/terms/",
    "contact": {
      "name": "Swagger API Team"
    },
    "license": {
      "name": "MIT"
    }
  },
  "host": "petstore.swagger.io",
  "basePath": "/api",
  "schemes": [
    "http"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/pets": {
      "get": {
        "description": "Returns all pets from the system that the user has access to",
        "produces": [
          "application/json"
        ],
        "responses": {
          "200": {
            "description": "A list of pets.",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Pet"
              }
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Pet": {
      "type": "object",
      "required": [
        "id",
        "name"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string"
        },
        "tag": {
          "type": "string"
        }
      }
    }
  }
}
`
