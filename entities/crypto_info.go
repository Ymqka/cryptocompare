package entities

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4"

	"github.com/Ymqka/cryptocompare/cache"
	"github.com/Ymqka/cryptocompare/repository"
)

const CryptoInfoCacheKey = "cryptoInfoKey"

type CryptoInfoI interface {
	GetLatestCryptoInfo(ctx context.Context) (latestInfo []byte, err error)
	AddCryptoInfo(ctx context.Context, pair string) (err error)
}

func NewCryptoInfoRepoCache(pgConn *pgx.Conn, redis *redis.Client) CryptoInfoI {
	return &cache.CryptoInfoRepoCache{
		CryptoInfoRepo: repository.CryptoInfoRepo{PgConn: pgConn},
		Redis:          redis,
	}
}
