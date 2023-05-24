package v1

import (
	context "context"
	fmt "fmt"
	"log"
	"net"

	"kratos-base/pkg/mock"
	grpc "google.golang.org/grpc"
)

// NewMockServer creates a new mock billing gRPC server
func NewMockServer(port int, messageChan chan interface{}) {
	// create a listener on passed TCP port
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a server instance
	s := MockServer{messageChan: messageChan}

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// attach the Ping service to the server
	RegisterCatalogServer(grpcServer, &s)

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

// MockServer represents the mocked gRPC server
type MockServer struct {
	messageChan chan interface{}
}

// CreateBeer 
func (m *MockServer) CreateBeer(ctx context.Context, req *CreateBeerReq) (*CreateBeerReply, error) {
	m.messageChan <- req

	response := mock.GetInterface(m.messageChan)
	if response == nil {
		log.Fatal("Test case did not program a mock response")
	}

	v, ok := response.(*CreateBeerReply)
	if !ok {
		return nil, fmt.Errorf("Mock response is the wrong type: %+v", response)
	}

	return v, nil
}

// GetBeer 
func (m *MockServer) GetBeer(ctx context.Context, req *GetBeerReq) (*GetBeerReply, error) {
	m.messageChan <- req

	response := mock.GetInterface(m.messageChan)
	if response == nil {
		log.Fatal("Test case did not program a mock response")
	}

	v, ok := response.(*GetBeerReply)
	if !ok {
		return nil, fmt.Errorf("Mock response is the wrong type: %+v", response)
	}

	return v, nil
}


func (m *MockServer) mustEmbedUnimplementedCatalogServer(){
	
}