# internal Directory

The `internal` directory contains the core business logic of the application. The code here is not intended to be imported by other applications, ensuring that the internal implementation details are encapsulated and protected.

## Directory Structure

The typical structure of the `internal` directory may include subdirectories for handlers, routers, services, models, and documentation.

### handler

The `handler` subdirectory contains the HTTP handlers that manage incoming requests and responses. Each handler function maps to a specific endpoint in your API.

#### Typical Files:

- `ping.go`: Contains the handler for the `/ping` endpoint.
- `hello.go`: Contains the handler for the `/hello` endpoint.

### router

The `router` subdirectory contains the routing logic that maps endpoints to their corresponding handlers. This is where you set up the routes and any necessary middleware.

#### Typical Files:

- `router.go`: Sets up the routes for the application.

### service

The `service` subdirectory contains the business logic of the application. Services are used by handlers to perform operations and should be independent of the HTTP layer.

#### Typical Files:

- `service.go`: Contains the core business logic for the application.

### model

The `model` subdirectory contains the data structures and database models used in the application. These models are used by services to interact with the data layer.

#### Typical Files:

- `model.go`: Defines the data structures and database models.

### docs

The `docs` subdirectory contains the Swagger documentation generated for the API. This directory is typically auto-generated and should not be manually edited.

#### Typical Files:

- `docs.go`: Contains the generated Go code for Swagger documentation.
- `swagger.json`: Contains the generated Swagger JSON documentation.
- `swagger.yaml`: Contains the generated Swagger YAML documentation.

## Code Examples

### handler/ping.go

```go
package handler

import (
    "github.com/gin-gonic/gin"
)

// PingEndpoint godoc
// @Summary Ping example
// @Description Do ping
// @Tags example
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /ping [get]
func PingEndpoint(c *gin.Context) {
    c.JSON(200, gin.H{"message": "pong"})
}