package container

import (
	"context"
	"fmt"

	"github.com/SeaOfWisdom/sow_contractor/src/config"
	"github.com/SeaOfWisdom/sow_contractor/src/log"
	"github.com/SeaOfWisdom/sow_contractor/src/server"
	srv "github.com/SeaOfWisdom/sow_contractor/src/service"
	"github.com/SeaOfWisdom/sow_contractor/src/storage"
	libProto "github.com/SeaOfWisdom/sow_proto/lib-srv"

	"github.com/go-redis/redis/v8"
	"github.com/olebedev/emitter"
	"go.uber.org/dig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateContainer() *dig.Container {
	container := dig.New()
	/* init base */
	must(container.Provide(log.NewLogger))
	must(container.Provide(config.NewConfig))
	must(container.Provide(func() *emitter.Emitter {
		return emitter.New(10)
	}))
	/* init redis */
	must(container.Provide(func(config *config.Config) *redis.Client {
		redisOpts, err := redis.ParseURL(fmt.Sprintf("redis://%s:%s@%s/", config.RedisUser, config.RedisPassword, config.RedisAddress))
		if err != nil {
			panic(fmt.Errorf("unable to parse redis URL: %v", err))
		}
		cli := redis.NewClient(redisOpts)
		if err := cli.Ping(context.Background()).Err(); err != nil {
			panic(fmt.Errorf("unable to set connection with redis: %v", err))
		}
		return cli
	}))
	/* initialize storage of redis database */
	must(container.Provide(storage.NewStorageService))
	/* set connection to the internal services */
	/* Library Service */
	must(container.Provide(func(config *config.Config) libProto.LibraryServiceClient {
		conn, err := grpc.Dial(config.LibServiceGrpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(fmt.Errorf("unable to start Lib Service GRPC client: %v", err))
		}
		return libProto.NewLibraryServiceClient(conn)
	}))
	/* initialize internal services */
	must(container.Provide(srv.NewService))
	must(container.Provide(server.NewGrpcServer))
	return container
}

func MustInvoke(container *dig.Container, function interface{}, opts ...dig.InvokeOption) {
	must(container.Invoke(function, opts...))
}

func must(err error) {
	if err != nil {
		panic(fmt.Sprintf("failed to initialize DI: %s", err))
	}
}
