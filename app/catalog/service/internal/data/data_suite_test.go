package data_test

import (
	"context"
	"kratos-base/app/catalog/service/internal/conf"
	"kratos-base/app/catalog/service/internal/data"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// 测试 data 方法
func TestData(t *testing.T) {
	//  Ginkgo 测试通过调用 Fail(description string) 功能来表示失败
	// 使用 RegisterFailHandler 将此函数传递给 Gomega 。这是 Ginkgo 和 Gomega 之间的唯一连接点
	RegisterFailHandler(Fail)
	// 通知 Ginkgo 启动测试套件。如果您的任何 specs 失败，Ginkgo 将自动使 testing.T 失败。
	RunSpecs(t, "biz data test catalog")
}
var cleaner func()      // 定义删除 pg 容器的回调函数
var Db *data.Data       // 用于测试的 data
var ctx context.Context // 上下文


// ginkgo 使用 BeforeEach 为您的 Specs 设置状态
var _ = BeforeSuite(func() {

	// 执行测试数据库操作之前，链接之前 docker 容器创建的 pg
	con, f := DockerPostgreSQL()
	cleaner = f
	config := &conf.Data{Database: &conf.Data_Database{Driver: "postgres", Source: con}}
	db, _ := data.NewEntClient(config, log.DefaultLogger)

	postgreSQLDB, _, err := data.NewData(config, db, log.DefaultLogger)
	if err != nil {
		return
	}
	Db = postgreSQLDB
	ctx = context.Background()
	Expect(err).NotTo(HaveOccurred())
})

// 测试结束后 通过回调函数，关闭并删除 docker 创建的容器
var _ = AfterSuite(func() {
	cleaner()
})
