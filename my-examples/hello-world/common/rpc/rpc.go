package rpc

import (
	"context"
	"encoding/json"
	"go-micro.dev/v5/client"
)

func GrpcCall(c client.Client, serviceName, methodName string, reqBody []byte) ([]byte, error) {

	requestData := json.RawMessage(reqBody)

	// 动态构造 RPC 请求
	req := c.NewRequest(
		serviceName,  // 服务名
		methodName,   // 方法名
		&requestData, // 传递指针
		//client.WithContentType("application/json"),
	)

	var rsp json.RawMessage
	err := c.Call(context.Background(), req, &rsp)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}
