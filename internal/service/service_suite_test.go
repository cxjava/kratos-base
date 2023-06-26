package service_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "service test")
}

var ctl *gomock.Controller
var cleaner func()
var ctx context.Context

var _ = BeforeEach(func() {
	ctl = gomock.NewController(GinkgoT())
	cleaner = ctl.Finish
	ctx = context.Background()
})

var _ = AfterEach(func() {
	// remove any mocks
	cleaner()
})
