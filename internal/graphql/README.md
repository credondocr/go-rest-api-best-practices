
# GraphQL Schema and Resolvers

This directory contains the GraphQL schema and resolver definitions for the Go REST API Best Practices project. The GraphQL layer connects to the gRPC service to fetch data.

## Directory Structure

- **schema.go**: Contains the GraphQL schema definition and resolver functions.

## Setup

Ensure that you have the necessary dependencies installed:

```sh
go get github.com/graphql-go/graphql
go get github.com/graphql-go/handler
go get google.golang.org/grpc
```

## GraphQL Schema

The GraphQL schema is defined in `schema.go`. It includes the following types and queries:

### Types

- **Product**: Represents a product with fields `id`, `name`, `description`, and `price`.

### Queries

- **product**: Fetches a product by its ID.

### Example Query

Fetch a product by ID:

```graphql
{
  product(id: "12345") {
    id
    name
    description
    price
  }
}
```

## Resolving Queries

The resolvers are responsible for fetching data from the gRPC service. The example below shows how the `product` query is resolved by calling the gRPC service.

### Example Resolver

```go
Resolve: func(p graphql.ResolveParams) (interface{}, error) {
    id, ok := p.Args["id"].(string)
    if !ok {
        return nil, nil
    }

    // Connect to gRPC service
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        return nil, err
    }
    defer conn.Close()

    client := pb.NewExampleServiceClient(conn)
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    req := &pb.ProductRequest{Id: id}
    res, err := client.GetProduct(ctx, req)
    if err != nil {
        return nil, err
    }

    return res, nil
}
```

## Running the GraphQL Server

Ensure your main server is set up to handle GraphQL requests:

### Example Configuration in `main.go`

```go
graphqlHandler := handler.New(&handler.Config{
    Schema:   &graphql.Schema,
    Pretty:   true,
    GraphiQL: true,
})
r.POST("/graphql", gin.WrapH(graphqlHandler))
r.GET("/graphql", gin.WrapH(graphqlHandler))
```

With this setup, you can access the GraphQL endpoint at `http://localhost:8080/graphql` and use tools like GraphiQL to run queries against your schema.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
