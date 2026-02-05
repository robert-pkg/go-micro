package main

import (
	"io"

	"github.com/gin-gonic/gin"
	micro "go-micro.dev/v5"
	"go-micro.dev/v5/client"
	"go-micro.dev/v5/client/grpc"
	logger "go-micro.dev/v5/logger"
	"go-micro.dev/v5/registry"
	"go-micro.dev/v5/registry/consul"

	log "github.com/robert-pkg/go-micro/my-examples/hello-world/common/log"
	"github.com/robert-pkg/go-micro/my-examples/hello-world/common/rpc"
)

func main() {
	if err := log.Init("", true); err != nil {
		panic(err)
	}

	logger.Info("start")

	service := micro.NewService(
		micro.Registry(consul.NewConsulRegistry(
			registry.Addrs("127.0.0.1:8500"), // consul 地址
		)),
		micro.Client(grpc.NewClient(
			client.ContentType("application/json"), // 设置内容类型
		)),
	)
	service.Init()

	r := gin.Default()
	// POST /api/:service/:method
	r.POST("/api/:service/:method", func(c *gin.Context) {
		serviceName := c.Param("service")
		methodName := serviceName + "." + c.Param("method")

		logger.Info("call", " serviceName=", serviceName, " methodName=", methodName)

		// 读取 body
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid request body"})
			return
		}

		bs, err := rpc.GrpcCall(service.Client(), serviceName, methodName, body)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// 返回响应
		c.Data(200, "application/json", bs)
	})

	r.Run(":9098")

}
