package graphql

import (
	"context"
	"time"

	"github.com/graphql-go/graphql"
	"google.golang.org/grpc"

	pb "github.com/credondocr/go-rest-api-best-practices/proto"
)

var ProductType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Product",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"price": &graphql.Field{
			Type: graphql.Float,
		},
	},
})

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"product": &graphql.Field{
			Type: ProductType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(string)
				if !ok {
					return nil, nil
				}

				conn, err := grpc.Dial("grpc-server:50051", grpc.WithInsecure())
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
			},
		},
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: RootQuery,
})
