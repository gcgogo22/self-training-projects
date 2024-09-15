package main

import (
	"context"
	"fmt"
	"log"
)

// Reach to the productpb package that is generated from protobuffer code.
func main() {
	p := productpb.Product{
		Id:          int32(products[0].ID),
		Name:        products[0].Name,
		UserPerUnit: products[0].USDPerUnit,
		Unit:        products[0].Unit,
	}

	// Need to convert this Protocol Buffer message into a format that can be sent over the network. Instead of using the JSON message, we will use the proto package. -> go get google.golang.org/protobuf

	data, err := proto.Marshal(&p)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data)) // check what's the output -> this gives not readable output.

	// How to convert it back
	var p2 productpb.Product
	err = proto.Unmarshal(data, &p2) // Like unmarshal message
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", p2) // Check extra informatioin.
}

// Define the proto buffer service
type ProductService struct {
	productpb.UnimplementedProductServer
}

func (ps ProductService) GetProduct(ctx context.Context, req *productpb.GetProductRequest) (*productpb.GetProductReply, error) {
	// Business logic
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

// Get Server start -> You can start the server either with tcp protocal or register it as one endpoint of the http server. Either one is supported. 


/*
To register the GetProduct service to the Server, you need to download a new package. 

go get google.golang.org/grpc
*/
func startGRPCServer() {
	lis, err := net.Listen("tcp", "localhost:4001")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	productpb.RegisterProductServer(grpcServer, &ProductService{})
	log.Fatal(grpcServer.Serve(lis))
}

// Create a client and call this service
func callGRPCService() {
	// get connection
	opts := []grpc.DialOption{grpc.WithInsecure()} // skip TLS
	conn, err := grpc.Dial("localhost:4001", opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// create client
	client := productpb.NewProductClient(conn)
	res, err := client.GetProduct(context.TODO(), &productpb.GetProductRequest{ProductId: 3})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", res.Product)
}

/*
Main function to start up the grpc service

func main() {
go startGRPCServer()

time.Sleep(1 * time.Second)

callGRPCService()

}
*/
