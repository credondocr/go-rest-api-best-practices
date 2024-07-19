package service

import (
	"context"
	"io"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	pb "github.com/credondocr/go-rest-api-best-practices/proto"
)

var grpcClient pb.ExampleServiceClient

func InitHandler(grpcAddr string) {
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	grpcClient = pb.NewExampleServiceClient(conn)

}

func CallGRPCService(name string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ExampleRequest{Name: name}
	res, err := grpcClient.GetExample(ctx, req)
	if err != nil {
		return "", err
	}

	return res.Message, nil
}

func GetAllProducts() ([]*pb.Product, error) {
	stream, err := grpcClient.GetAllProducts(context.Background(), &pb.Empty{})
	if err != nil {
		return nil, err
	}

	var products []*pb.Product
	for {
		product, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func SearchProducts(name string, omitCache bool) ([]*pb.Product, error) {

	ctx := context.Background()

	if omitCache {
		md := metadata.Pairs("X-Omit-Cache", "true")
		ctx = metadata.NewOutgoingContext(ctx, md)
	}

	stream, err := grpcClient.SearchProducts(ctx, &pb.SearchRequest{Name: name})
	if err != nil {
		return nil, err
	}

	var products []*pb.Product
	for {
		product, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
