package storage

import (
	"context"
	"fmt"

	"github.com/SeaOfWisdom/sow_contractor/src/config"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

const (
	latestBlockKey = "latestBlock"
)

type StorageService struct {
	startBlock uint64
	redisCli   *redis.Client
}

func NewStorageService(cfg *config.Config, rClient *redis.Client) *StorageService {
	return &StorageService{
		redisCli:   rClient,
		startBlock: cfg.StartFromBlock,
	}
}

/*///////////////////////
///// Latest block /////
/////////////////////*/

func (s *StorageService) GetLatestKnownBlock() (uint64, error) {
	latestBlock, err := s.redisCli.Get(context.Background(), latestBlockKey).Uint64()
	if err == redis.Nil {
		return s.startBlock, nil
	} else if err != nil {
		return 0, fmt.Errorf("unable to update latest block, err: %v", err)
	}
	return latestBlock, nil
}

func (s *StorageService) UpdateLatestKnownBlock(latestKnownBlock uint64) error {
	if err := s.redisCli.Set(context.Background(), latestBlockKey, latestKnownBlock, redis.KeepTTL).Err(); err != nil {
		return errors.Wrap(err, "unable to update latest block")
	}

	return nil
}
