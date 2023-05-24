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

	userv1 "kratos-base/api/user/service/v1"
	"kratos-base/pkg/test_util"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpPort = 50051
	grpcPort = 50053
	catalogPort = 50054
	deferFunc = []func ()  {
		
	}
)

var (
	userClient userv1.UserClient
)


func startCatalogService() {

	httpPort = test_util.OpenPort()
	catalogPort = test_util.OpenPort()

	con, dbPurge := DockerPostgreSQL()

	log.Println("start catalog test server")
	err := os.Chdir("../../app/catalog/service/cmd/server/")
	if err != nil {
		log.Fatal(err)
		return
	}

	cmd := exec.Command("go", "run", ".")
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "KRATOS_DB_SOURCE="+con)
	cmd.Env = append(cmd.Env, "KRATOS_HTTP_ADDR=localhost:"+strconv.Itoa(httpPort))
	cmd.Env = append(cmd.Env, "KRATOS_GRPC_ADDR=localhost:"+strconv.Itoa(catalogPort))
	go cmd.Start()

	// wait some time until mock services are started
	up1 := make(chan struct{})
	go test_util.CheckServerIsReady(catalogPort, up1)
	<-up1
	deferFunc = append(deferFunc, dbPurge)
	deferFunc = append(deferFunc, func() {cmd.Process.Kill()})
}
func startUserService() {

	httpPort = test_util.OpenPort()
	grpcPort = test_util.OpenPort()

	con, dbPurge := DockerPostgreSQL()

	log.Println("start user test server")
	// take care this path
	err := os.Chdir("../../../../user/service/cmd/server/")
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

	// wait some time until mock services are started
	up1 := make(chan struct{})
	go test_util.CheckServerIsReady(grpcPort, up1)
	<-up1
	deferFunc = append(deferFunc, dbPurge)
	deferFunc = append(deferFunc, func() {cmd.Process.Kill()})
}
func TestMain(m *testing.M) {

	startCatalogService()
	startUserService()

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
	for _, f := range deferFunc {
		f()
	}
	os.Exit(code)
}

func TestCreateUser(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	//save first
	userClient.CreateUser(ctx, &userv1.UserCreateReq{Age: 1, Name: "test2"})

	r, err := userClient.CreateUser(ctx, &userv1.UserCreateReq{Age: 1, Name: "test1"})
	if err != nil {
		t.Fatalf("CreateUser error: %v", err)
	}
	assert.Equal(t, "success", r.Message, "create user should be success")
	assert.Equal(t, "test1_test1_Beer", r.Result.Name, "user's name should be correct")
	assert.Equal(t, int64(2), r.Result.Id, "user's id should be 2")
	
}
