package functions

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"hasura-ndc.dev/ndc-go/types"
)

// ServiceInfoResponse represents the structure of the REST API response
type ServiceInfoResponse struct {
	Data struct {
		ServiceInfo struct {
			CurrentVersion string `json:"currentVersion"`
			Environment    string `json:"environment"`
			UpdatedAt      string `json:"updatedAt"`
		} `json:"serviceInfo"`
	} `json:"data"`
}

// DatabaseRow represents the data returned from the database query
type DatabaseRow struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ToMap method is provided by types.generated.go

// BusinessLogicArguments defines the arguments for the BusinessLogic function
type BusinessLogicArguments struct {
	// You can add any arguments here if needed
}

// FromValue decodes values from map
func (j *BusinessLogicArguments) FromValue(input map[string]any) error {
	// No arguments to parse for now, but we need this method for the connector
	return nil
}

// BusinessLogicResult defines the result of the BusinessLogic function
type BusinessLogicResult struct {
	Environment string      `json:"environment"`
	RowData     DatabaseRow `json:"rowData"`
}

// Note: ToMap method is now provided by types.generated.go

// FunctionBusinessLogic calls a REST API, extracts the environment value,
// and uses it to query a Postgres database
//
// Example:
//
//	curl http://localhost:8080/query -H 'content-type: application/json' -d \
//		'{
//			"collection": "businessLogic",
//			"arguments": {},
//			"collection_relationships": {},
//			"query": {
//				"fields": {
//					"environment": {
//						"type": "column",
//						"column": "environment"
//					},
//					"rowData": {
//						"type": "column",
//						"column": "rowData"
//					}
//				}
//			}
//		}'
func FunctionBusinessLogic(ctx context.Context, state *types.State, arguments *BusinessLogicArguments) (*BusinessLogicResult, error) {
	// Load environment variables from .env file (ignore error if file doesn't exist)
	_ = godotenv.Load()

	// Call the REST API
	environment, err := getEnvironmentFromAPI()
	if err != nil {
		return nil, fmt.Errorf("error getting environment from API: %w", err)
	}

	// Query the database using the environment value
	rowData, err := queryDatabaseByEnvironment(environment)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %w", err)
	}

	// Return the result
	return &BusinessLogicResult{
		Environment: environment,
		RowData:     rowData,
	}, nil
}

// getEnvironmentFromAPI calls the REST API and extracts the environment value
func getEnvironmentFromAPI() (string, error) {
	// In a real implementation, this would make an HTTP request
	// For now, we'll return a hardcoded response that matches the expected format
	jsonResponse := `{
		"data": {
			"serviceInfo": {
				"currentVersion": "4.0.100",
				"environment": "stage",
				"updatedAt": "2025-05-14T18:10:02.387556813Z"
			}
		}
	}`

	// Parse the JSON response
	var serviceInfo ServiceInfoResponse
	if err := json.Unmarshal([]byte(jsonResponse), &serviceInfo); err != nil {
		return "", fmt.Errorf("error parsing JSON response: %w", err)
	}

	// Extract and return the environment value
	return serviceInfo.Data.ServiceInfo.Environment, nil
}

// queryDatabaseByEnvironment queries the Postgres database using the environment value
func queryDatabaseByEnvironment(environment string) (DatabaseRow, error) {
	// Check if we're in development/testing mode
	if os.Getenv("APP_ENV") == "development" || os.Getenv("DB_HOST") == "" {
		// Return mock data for development/testing
		return DatabaseRow{
			ID:          "1",
			Name:        fmt.Sprintf("%s-service", environment),
			Description: fmt.Sprintf("This is a %s environment service", environment),
		}, nil
	}

	// Get database connection parameters from environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Construct the connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return DatabaseRow{}, fmt.Errorf("error connecting to database: %w", err)
	}
	defer db.Close()

	// Query the database
	// query := "SELECT id, name, description FROM table_name WHERE environment = $1 LIMIT 1"
	query := "SELECT opportunity_id as id, amount as name, stage as description FROM opportunities WHERE opportunity_id = 'O3000' LIMIT 1"

	//row := db.QueryRow(query, environment)
	row := db.QueryRow(query)

	// Scan the result into a DatabaseRow struct
	var result DatabaseRow
	err = row.Scan(&result.ID, &result.Name, &result.Description)
	if err != nil {
		return DatabaseRow{}, fmt.Errorf("error scanning database row: %w", err)
	}

	return result, nil
}
