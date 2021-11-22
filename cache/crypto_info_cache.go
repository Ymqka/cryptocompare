package cache

import (
	"context"

	"github.com/Ymqka/cryptocompare/repository"

	"github.com/go-redis/redis/v8"
)

const CryptoInfoCacheKey = "cryptoInfoKey"

type CryptoInfoRepoCache struct {
	repository.CryptoInfoRepo
	Redis *redis.Client
}

func (mr *CryptoInfoRepoCache) GetLatestCryptoInfo(ctx context.Context) (latestInfo []byte, err error) {
	var result string
	if result, err = mr.Redis.Get(ctx, CryptoInfoCacheKey).Result(); err != nil {
		return nil, err
	}

	return []byte(result), nil
}
