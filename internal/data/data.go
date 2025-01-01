package data

import (
	"context"
	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"
	v1 "review-b/api/review/v1"
	"review-b/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewBusinessRepo, NewReviewServiceClient, NewDiscovery)

// Data .
type Data struct {
	// TODO wrapped database client
	// 嵌入一个gRPC client，通过这个client调用review-service服务
	rc  v1.ReviewClient
	log *log.Helper
}

// NewData .
func NewData(c *conf.Data, rc v1.ReviewClient, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		rc:  rc,
		log: log.NewHelper(logger),
	}, cleanup, nil
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	config := api.DefaultConfig()
	config.Address = conf.Consul.Address
	config.Scheme = conf.Consul.Scheme
	// new consul client
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	// new dis with consul client
	dis := consul.New(client)
	return dis
}

func NewReviewServiceClient(d registry.Discovery) v1.ReviewClient {
	// 这里需要导入kratos包装好的grpc包
	// import "github.com/go-kratos/kratos/v2/transport/grpc"
	conn, err := grpc.DialInsecure(
		context.Background(),
		//grpc.WithEndpoint("127.0.0.1:9000"),
		grpc.WithEndpoint("discovery:///review.service"),
		grpc.WithDiscovery(d),
		// 还可以携带中间件
		grpc.WithMiddleware(
			recovery.Recovery(),
			validate.Validator(),
		),
	)
	if err != nil {
		panic(err)
	}
	return v1.NewReviewClient(conn)
}
