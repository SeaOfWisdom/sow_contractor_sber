package container

import (
	"fmt"
	"time"

	"github.com/SeaOfWisdom/sow_contractor/src/config"
	"github.com/SeaOfWisdom/sow_contractor/src/server"
	srv "github.com/SeaOfWisdom/sow_contractor/src/service"
	"github.com/olebedev/emitter"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

func CreateContainer() *dig.Container {
	container := dig.New()
	/* init base */
	must(container.Provide(NewLogger))
	must(container.Provide(config.NewConfig))
	must(container.Provide(func() *emitter.Emitter {
		return emitter.New(10)
	}))
	/* set connection to the internal services */
	/* initialize internal services */
	must(container.Provide(srv.NewService))
	must(container.Provide(server.NewGrpcServer))
	return container
}

func NewLogger() *logrus.Logger {
	// init logrus logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC822,
	})
	// set logger level
	level, err := logrus.ParseLevel("debug")
	if err != nil {
		panic(err)
	}
	logger.SetLevel(level)
	return logger
}

func MustInvoke(container *dig.Container, function interface{}, opts ...dig.InvokeOption) {
	must(container.Invoke(function, opts...))
}

func must(err error) {
	if err != nil {
		panic(fmt.Sprintf("failed to initialize DI: %s", err))
	}
}
