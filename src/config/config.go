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
	BscProvider           string
	BscChainID            int64
	BscOperatorPrivateKey string
	BscOperatorAddress    string
	BscCheckerSleepTime   time.Duration
	BscWaitForBlocks      uint64
	/* Smart contracts */

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
	flag.StringVar(&config.BscProvider, "bsc-provider", "https://data-seed-prebsc-2-s2.binance.org:8545", "")
	flag.Int64Var(&config.BscChainID, "bsc-chain-id", 97, "")
	flag.StringVar(&config.BscOperatorPrivateKey, "bsc-operator-private-key", "", "")
	flag.StringVar(&config.BscOperatorAddress, "bsc-operator-address", "0xdD868980eF73eDCbC1fF758F6e53023bE18e2A52", "")
	flag.Uint64Var(&config.BscWaitForBlocks, "bsc-wait-for-blocks", 10, "")
	flag.DurationVar(&config.BscCheckerSleepTime, "bsc-checker-sleep-time", 3*time.Second, "")
	/* Smart contracts */

	/* parse config from envs or config files */
	flag.Parse()
	return config
}
