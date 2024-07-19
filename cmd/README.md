# cmd Directory

The `cmd` directory contains the entry points for the application. Each subdirectory in `cmd` represents a different executable or service that your project provides. This separation allows you to organize multiple command-line applications or services within the same repository.

## Directory Structure

The typical structure of the `cmd` directory may include subdirectories for each executable. In this example, we have a single `server` subdirectory for the main API server.

### server

The `server` subdirectory contains the main entry point for the API server. This is where the application is initialized and the server is started.

#### Typical Files:

- `main.go`: The main entry point for the API server.

## Code Example

### server/main.go

```go
package main

import (
    "myapi/internal/router"
    "myapi/pkg/config"
    _ "myapi/internal/docs" // Import Swagger documentation

    "github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
)

// @title My API
// @version 1.0
// @description This is a sample server.
// @host localhost:8080
// @BasePath /api/v1

func main() {
    cfg := config.LoadConfig()
    r := router.SetupRouter()

    // Swagger
    r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

    r.Run(cfg.Server.Port)
}