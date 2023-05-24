package test

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"testing"
	"time"

	catalogv1 "kratos-base/api/catalog/service/v1"
	userv1 "kratos-base/api/user/service/v1"
	"kratos-base/pkg/mock"
	"kratos-base/pkg/test_util"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpPort = 50051
	grpcPort = 50053
	catalogPort = 50054
)

var (
	userClient userv1.UserClient
	catalogChannel  chan interface{}
)

func TestMain(m *testing.M) {

	httpPort = test_util.OpenPort()
	grpcPort = test_util.OpenPort()
	catalogPort = test_util.OpenPort()

	con, dbPurge := DockerPostgreSQL()

	log.Println("start integration test server")
	err := os.Chdir("../../cmd/server/")
	if err != nil {
		log.Fatal(err)
		return
	}

	cmd := exec.Command("go", "run", ".")
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "KRATOS_DB_SOURCE="+con)
	cmd.Env = append(cmd.Env, "KRATOS_HTTP_ADDR=localhost:"+strconv.Itoa(httpPort))
	cmd.Env = append(cmd.Env, "KRATOS_GRPC_ADDR=localhost:"+strconv.Itoa(grpcPort))
	cmd.Env = append(cmd.Env, "KRATOS_CATALOG_ADDR=localhost:"+strconv.Itoa(catalogPort))
	go cmd.Start()


	// Run downstream mocks
	catalogChannel = make(chan interface{}, 100)
	go catalogv1.NewMockServer(catalogPort, catalogChannel)

	// wait some time until mock services are started
	up1 := make(chan struct{})
	up2 := make(chan struct{})
	go test_util.CheckServerIsReady(grpcPort, up1)
	go test_util.CheckServerIsReady(catalogPort, up2)
	<-up1
	<-up2

	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", grpcPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	userClient = userv1.NewUserClient(conn)

	// Run tests
	code := m.Run()
	// clean up
	// You can't defer this because os.Exit doesn't care for defer
	dbPurge()
	cmd.Process.Kill()

	os.Exit(code)
}

func TestCreateUser(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	catalogChannel <- &catalogv1.CreateBeerReply{Id: 1, Name: "From Mock"}
	
	r, err := userClient.CreateUser(ctx, &userv1.UserCreateReq{Age: 1, Name: "test1"})
	if err != nil {
		t.Fatalf("CreateUser error: %v", err)
	}
	assert.Equal(t, "success", r.Message, "create user should be success")
	assert.Equal(t, "test1_From Mock", r.Result.Name, "user's name should be correct")

	// Validate Accounts Server made a request to the Payments Service
	request := mock.GetInterface(catalogChannel)
	if request == nil {
		t.Fatal("User service did not make a request to Catalog Service")
	}
	pr, ok := request.(*catalogv1.CreateBeerReq)
	if !ok {
		t.Fatalf("Catalog Service recieved the wrong type of request: %v", pr)
	}
	assert.Equal(t, "test1_Beer", pr.Name, "check Catalog Service recieved correct Name")

}
