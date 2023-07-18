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
	/* MongoDB */
	MongoAddress  string
	MongoUser     string
	MongoPassword string
	MongoDBName   string
	/* BSC Blockchain */
	Provider           string
	ChainID            int64
	OperatorPrivateKey string
	OperatorAddress    string
	CheckerSleepTime   time.Duration
	WaitForBlocks      uint64
	/* Smart contracts */
	SowTokenAddress   string
	SowLibraryAddress string

	/* Internal communication services */
	JWTServiceGRpcAddress string
	OCRServiceGRpcAddress string
}

func NewConfig() *Config {
	config := &Config{}
	/* gRPC */
	flag.StringVar(&config.GrpcAddress, "grpc-address", "0.0.0.0:5305", "gRPC address and port for inter-service communications")
	/* Prometheus */
	flag.StringVar(&config.PrometheusAddress, "prometheus-address", "localhost:8075", "host and port for prometheus")
	/* BSC Blockchain */
	flag.StringVar(&config.Provider, "provider", "https://rpc.test.siberium.net", "")
	flag.Int64Var(&config.ChainID, "bsc-chain-id", 111000, "")
	flag.StringVar(&config.OperatorPrivateKey, "bsc-operator-private-key", "", "")
	flag.StringVar(&config.OperatorAddress, "bsc-operator-address", "0xdD868980eF73eDCbC1fF758F6e53023bE18e2A52", "")
	flag.Uint64Var(&config.WaitForBlocks, "bsc-wait-for-blocks", 10, "")
	flag.DurationVar(&config.CheckerSleepTime, "bsc-checker-sleep-time", 3*time.Second, "")
	/* Smart contracts */
	flag.StringVar(&config.SowTokenAddress, "sow-token-address", "", "")
	flag.StringVar(&config.SowLibraryAddress, "sow-library-address", "", "")

	/* parse config from envs or config files */
	flag.Parse()
	return config
}
