// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package main

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/contrib/registry/consul/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"

	"main/grpc/protobuf"
)

func main() {
	registry, err := consul.New(consul.WithAddress("127.0.0.1:8500"))
	if err != nil {
		g.Log().Fatal(context.Background(), err)
	}
	grpcx.Resolver.Register(registry)

	var (
		ctx    = context.Background()
		conn   = grpcx.Client.MustNewGrpcClientConn("demo")
		client = protobuf.NewGreeterClient(conn)
	)
	res, err := client.SayHello(ctx, &protobuf.HelloRequest{Name: "World"})
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	g.Log().Debug(ctx, "Response:", res.Message)
}
