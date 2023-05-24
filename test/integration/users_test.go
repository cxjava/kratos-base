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

	userv1 "kratos-base/api/user/v1"
	"kratos-base/pkg/test_util"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpPort = 50051
	grpcPort = 50053
)

var (
	userClient userv1.UserClient
)

func TestMain(m *testing.M) {

	httpPort = test_util.OpenPort()
	grpcPort = test_util.OpenPort()

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
	go cmd.Start()

	// wait some time until mock services are started
	up1 := make(chan struct{})
	go test_util.CheckServerIsReady(grpcPort, up1)
	<-up1

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
	r, err := userClient.CreateUser(ctx, &userv1.UserCreateReq{Age: 1, Name: "test1"})
	if err != nil {
		t.Fatalf("CreateUser error: %v", err)
	}
	assert.Equal(t, "success", r.Message, "create user should be success")
}
