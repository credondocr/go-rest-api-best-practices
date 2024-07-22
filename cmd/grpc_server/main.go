package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-redis/redis"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/credondocr/go-rest-api-best-practices/internal/model"
	pb "github.com/credondocr/go-rest-api-best-practices/proto"
)

var rdb *redis.Client

type server struct {
	pb.UnimplementedExampleServiceServer
	db *gorm.DB
}

func (s *server) GetExample(ctx context.Context, req *pb.ExampleRequest) (*pb.ExampleResponse, error) {
	return &pb.ExampleResponse{Message: "Hello " + req.Name}, nil
}

func (s *server) GetProduct(ctx context.Context, req *pb.ProductRequest) (*pb.Product, error) {
	var product model.Product
	result := s.db.First(&product, "id = ?", req.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &pb.Product{
		Id:          product.ID.String(),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}, nil
}

func (s *server) GetAllProducts(req *pb.Empty, stream pb.ExampleService_GetAllProductsServer) error {
	ctx := stream.Context()
	omitCache := metautils.ExtractIncoming(ctx).Get("X-Omit-Cache")

	if omitCache != "true" {
		cachedProducts, err := rdb.Get("products").Result()
		if err == redis.Nil {
			var products []model.Product
			result := s.db.Find(&products)
			if result.Error != nil {
				return result.Error
			}

			productsJSON, err := json.Marshal(products)
			if err != nil {
				return err
			}

			err = rdb.Set("products", productsJSON, 5*time.Minute).Err()
			if err != nil {
				return err
			}

			for _, product := range products {
				prod := &pb.Product{
					Id:          product.ID.String(),
					Name:        product.Name,
					Description: product.Description,
					Price:       product.Price,
				}
				if err := stream.Send(prod); err != nil {
					return err
				}
			}
		} else if err != nil {
			return err
		} else {
			var products []model.Product
			err := json.Unmarshal([]byte(cachedProducts), &products)
			if err != nil {
				return err
			}

			for _, product := range products {
				prod := &pb.Product{
					Id:          product.ID.String(),
					Name:        product.Name,
					Description: product.Description,
					Price:       product.Price,
				}
				if err := stream.Send(prod); err != nil {
					return err
				}
			}
		}
	} else {
		var products []model.Product
		result := s.db.Find(&products)
		if result.Error != nil {
			return result.Error
		}

		for _, product := range products {
			prod := &pb.Product{
				Id:          product.ID.String(),
				Name:        product.Name,
				Description: product.Description,
				Price:       product.Price,
			}
			if err := stream.Send(prod); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *server) SearchProducts(req *pb.SearchRequest, stream pb.ExampleService_SearchProductsServer) error {
	ctx := stream.Context()
	name := req.GetName()
	cacheKey := "search:" + name
	omitCache := metautils.ExtractIncoming(ctx).Get("X-Omit-Cache")

	if omitCache != "true" {
		cachedProducts, err := rdb.Get(cacheKey).Result()
		if err == redis.Nil {
			var products []model.Product
			result := s.db.Where("name ILIKE ?", "%"+name+"%").Find(&products)
			if result.Error != nil {
				return result.Error
			}

			productsJSON, err := json.Marshal(products)
			if err != nil {
				return err
			}

			err = rdb.Set(cacheKey, productsJSON, 5*time.Minute).Err()
			if err != nil {
				return err
			}

			for _, product := range products {
				prod := &pb.Product{
					Id:          product.ID.String(),
					Name:        product.Name,
					Description: product.Description,
					Price:       product.Price,
				}
				if err := stream.Send(prod); err != nil {
					return err
				}
			}
		} else if err != nil {
			return err
		} else {
			var products []model.Product
			err := json.Unmarshal([]byte(cachedProducts), &products)
			if err != nil {
				return err
			}

			for _, product := range products {
				prod := &pb.Product{
					Id:          product.ID.String(),
					Name:        product.Name,
					Description: product.Description,
					Price:       product.Price,
				}
				if err := stream.Send(prod); err != nil {
					return err
				}
			}
		}
	} else {
		fmt.Println("skip cache")
		var products []model.Product
		result := s.db.Where("name ILIKE ?", "%"+name+"%").Find(&products)
		if result.Error != nil {
			return result.Error
		}

		for _, product := range products {
			prod := &pb.Product{
				Id:          product.ID.String(),
				Name:        product.Name,
				Description: product.Description,
				Price:       product.Price,
			}
			if err := stream.Send(prod); err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	dsn := "host=postgres user=user password=password dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&model.Product{})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)

	pb.RegisterExampleServiceServer(grpcServer, &server{db: db})

	grpc_prometheus.Register(grpcServer)
	grpc_prometheus.EnableHandlingTimeHistogram()

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Fatal(http.ListenAndServe(":9091", nil))
	}()

	log.Printf("gRPC server is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
