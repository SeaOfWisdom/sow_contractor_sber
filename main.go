package main

import (
	"github.com/SeaOfWisdom/sow_contractor/src/common"
	"github.com/SeaOfWisdom/sow_contractor/src/config"
	"github.com/SeaOfWisdom/sow_contractor/src/container"
	"github.com/SeaOfWisdom/sow_contractor/src/server"
	srv "github.com/SeaOfWisdom/sow_contractor/src/service"
)

func main() {
	di := container.CreateContainer()
	container.MustInvoke(di, func(
		config *config.Config,
		service *srv.Service,
		grpcServer *server.GrpcServer,
	) {
		/* start services */
		grpcServer.Start()
		service.Start()

		/* wait for application termination */
		common.WaitForSignal()

		grpcServer.Stop()
	})
}
