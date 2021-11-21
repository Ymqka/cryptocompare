package main

import (
	"context"
	"time"

	"cryptocompare/entities"

	"go.uber.org/zap"
)

// ParseCryptoCompare parses min-api.cryptocompare.com saves it to db then sleep for 60 seconds
func (rh *RequestHandler) ParseCryptoCompare(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		rh.parseCrypto(ctx)
		time.Sleep(60 * time.Second)
	}
}

// parseCrypto makes api request to cryptocompare and saves it to db
// if error occurs, logs it and trying to make request later
func (rh *RequestHandler) parseCrypto(ctx context.Context) {
	var (
		cryptoInfo []byte
		err        error
	)
	// requests api
	if cryptoInfo, err = rh.getCryptoInfo(
		ctx,
		rh.getFsymsFromConfig(),
		rh.getTsymsFromConfig(),
	); err != nil {
		rh.l.Error("failed to get info from cryptocompare", zap.Error(err))
		return
	}

	// saves to db
	if err = rh.cryptoInfo.AddCryptoInfo(ctx, string(cryptoInfo)); err != nil {
		rh.l.Error("failed to add pair to db", zap.Error(err))
		return
	}

	// saves to cache, ignore error
	_, _ = rh.cache.SetNX(ctx, entities.CryptoInfoCacheKey, string(cryptoInfo), 60*time.Second).Result()
}
