# pkg Directory

The `pkg` directory contains reusable packages that can be utilized by your application as well as other projects. This directory houses generic components and utilities that do not depend directly on the specific business logic of the application.

## Directory Structure

The typical structure of the `pkg` directory may include subdirectories for middleware, configuration, utilities, among others.

### middleware

The `middleware` subdirectory contains middleware for the Gin framework. Middleware are functions that are executed before or after the main request handling logic. These can be used for tasks such as authentication, logging, error handling, etc.

#### Typical Files:

- `middleware.go`: Contains the definition of common middleware used in the application.

### config

The `config` subdirectory contains logic related to loading and managing the application's configuration. It uses Viper to read configurations from JSON, YAML files, or environment variables.

#### Typical Files:

- `config.go`: Loads and manages the application's configuration.

### util

The `util` subdirectory contains generic utility functions that can be reused across different parts of the application. These utilities do not depend on business logic and may include helper functions for string manipulation, error handling, etc.

#### Typical Files:

- `util.go`: Contains generic utility functions.

## Code Examples

### middleware/middleware.go

```go
package middleware

import (
    "github.com/gin-gonic/gin"
    "log"
    "time"
)

func RequestLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        t := time.Now()

        c.Next()

        latency := time.Since(t)
        status := c.Writer.Status()

        log.Printf("Request: %s %s - %d - %v", c.Request.Method, c.Request.URL, status, latency)
    }
}