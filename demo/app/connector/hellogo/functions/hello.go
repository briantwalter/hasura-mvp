package functions

import (
	"context"
	"fmt"

	"hasura-ndc.dev/ndc-go/types"
)

// The generator infers the argument type in FunctionHello to generate codes
// Encoder and decoder methods will be generated to types.generated.go and connector.generated.go

// A hello argument
type HelloArguments struct {
	Greeting string `json:"greeting"` // value argument will be required
	Count    *int   `json:"count"`    // pointer arguments are optional
}

// A hello result
type HelloResult struct {
	Reply string `json:"reply"`
	Count int    `json:"count"`
}

// FunctionHello sends a hello message
// Function is an operation type of query, the name of function will be `hello`
//
// Example:
//
//	curl http://localhost:8080/query -H 'content-type: application/json' -d \
//		'{
//			"collection": "hello",
//			"arguments": {
//				"greeting": {
//					"type": "literal",
//					"value": "Hello world!"
//				}
//			},
//			"collection_relationships": {},
//			"query": {
//				"fields": {
//					"reply": {
//						"type": "column",
//						"column": "reply"
//					},
//					"count": {
//						"type": "column",
//						"column": "count"
//					}
//				}
//			}
//		}'
func FunctionHello(ctx context.Context, state *types.State, arguments *HelloArguments) (*HelloResult, error) {
	count := 1
	if arguments.Count != nil {
		count = *arguments.Count + 1
	}
	return &HelloResult{
		Reply: fmt.Sprintf("Hi! %s", arguments.Greeting),
		Count: count,
	}, nil
}

// Procedure is similar to function. However the data structure of arguments are simpler and can be decode directly with JSON

// A create author argument
type CreateAuthorArguments struct {
	Name string `json:"name"`
}

// A create author result
type CreateAuthorResult struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ProcedureCreateAuthor creates an author
// Procedure is an operation type of mutation, the name of function will be `createAuthor`
//
// Example:
//
//	curl http://localhost:8080/mutation -H 'content-type: application/json' -d \
//		'{
//		  "operations": [
//		    {
//		      "type": "procedure",
//		      "name": "createAuthor",
//		      "arguments": {
//						"name": "John"
//					},
//		      "fields": {
//		        "type": "object",
//		        "fields": {
//		          "id": {
//		            "type": "column",
//		            "column": "id"
//		          },
//		          "name": {
//		            "type": "column",
//		            "column": "name"
//		          }
//		        }
//		      }
//		    }
//		  ],
//		  "collection_relationships": {}
//		}'
func ProcedureCreateAuthor(ctx context.Context, state *types.State, arguments *CreateAuthorArguments) (*CreateAuthorResult, error) {
	return &CreateAuthorResult{
		ID:   1,
		Name: arguments.Name,
	}, nil
}
