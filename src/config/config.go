package config

import (
	"time"

	"github.com/namsral/flag"
)

type Config struct {
	/* gRPC */
	GrpcAddress string
	/* Prometheus */
	PrometheusAddress string
	/* Redis */
	RedisAddress  string
	RedisUser     string
	RedisPassword string
	/* BSC Blockchain */
	Provider           string
	ChainID            int64
	OperatorPrivateKey string
	OperatorAddress    string
	CheckerSleepTime   time.Duration
	StartFromBlock     uint64
	WaitForBlocks      uint64
	/* Smart contracts */
	SowTokenAddress   string
	SowLibraryAddress string
	/* Internal communication services */
	LibServiceGrpcAddress string
}

func NewConfig() *Config {
	config := &Config{}
	/* gRPC */
	flag.StringVar(&config.GrpcAddress, "grpc-address", "0.0.0.0:5305", "gRPC address and port for inter-service communications")
	/* Prometheus */
	flag.StringVar(&config.PrometheusAddress, "prometheus-address", "localhost:8075", "host and port for prometheus")
	/* Redis */
	flag.StringVar(&config.RedisAddress, "redis-address", "localhost:6379", "")
	flag.StringVar(&config.RedisUser, "redis-user", "", "")
	flag.StringVar(&config.RedisPassword, "redis-password", "", "")
	/* BSC Blockchain */
	flag.StringVar(&config.Provider, "provider", "https://rpc.test.siberium.net", "")
	flag.Int64Var(&config.ChainID, "chain-id", 111000, "")
	flag.StringVar(&config.OperatorPrivateKey, "operator-private-key", "", "")
	flag.StringVar(&config.OperatorAddress, "operator-address", "0xdD868980eF73eDCbC1fF758F6e53023bE18e2A52", "")
	flag.Uint64Var(&config.WaitForBlocks, "wait-for-blocks", 10, "")
	flag.Uint64Var(&config.StartFromBlock, "start-from-block", 599098, "")
	flag.DurationVar(&config.CheckerSleepTime, "checker-sleep-time", 3*time.Second, "")
	/* Smart contracts */
	flag.StringVar(&config.SowTokenAddress, "sow-token-address", "", "")
	flag.StringVar(&config.SowLibraryAddress, "sow-library-address", "0x56EAC6230bF6Cf9e674B2950F075945b662a9F69", "")
	/* Internal communication services */
	flag.StringVar(&config.LibServiceGrpcAddress, "lib-service-address", "0.0.0.0:8060", "")
	/* parse config from envs or config files */
	flag.Parse()
	return config
}
