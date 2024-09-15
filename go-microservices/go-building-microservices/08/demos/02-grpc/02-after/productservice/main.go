package main

import (
	"context"
	"demo/productpb"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type Product struct {
	ID         int
	Name       string
	USDPerUnit float64
	Unit       string
}

func main() {

	go startGRPCServer()

	time.Sleep(1 * time.Second)

	callGRPCService()

}

type ProductService struct {
	productpb.UnimplementedProductServer
}

func (ps ProductService) GetProduct(ctx context.Context, req *productpb.GetProductRequest) (*productpb.GetProductReply, error) {
	for _, p := range products {
		if p.ID == int(req.ProductId) {
			return &productpb.GetProductReply{
					Product: &productpb.Product{
						Id:         int32(p.ID),
						Name:       p.Name,
						UsdPerUnit: p.USDPerUnit,
						Unit:       p.Unit,
					},
				},
				nil
		}
	}

	return nil, fmt.Errorf("product not found with ID: %v", req.ProductId)
}

func startGRPCServer() {
	lis, err := net.Listen("tcp", "localhost:4001")
	if err != nil {
		log.Fatal(err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	productpb.RegisterProductServer(grpcServer, &ProductService{})
	log.Fatal(grpcServer.Serve(lis))
}

func callGRPCService() {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial("localhost:4001", opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := productpb.NewProductClient(conn)
	res, err := client.GetProduct(context.TODO(), &productpb.GetProductRequest{ProductId: 3})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Product)
}
