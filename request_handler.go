package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"cryptocompare/entities"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type RequestHandler struct {
	l          *zap.Logger
	cfg        *viper.Viper
	cache      *redis.Client
	cryptoInfo entities.CryptoInfoI
}

func NewRequestHandler(
	l *zap.Logger,
	cfg *viper.Viper,
	cryptoInfo entities.CryptoInfoI,
	cache *redis.Client,
) (rh *RequestHandler) {
	return &RequestHandler{
		l:          l,
		cfg:        cfg,
		cryptoInfo: cryptoInfo,
		cache:      cache,
	}
}

// init logger, config, pg connections
func initRequestHandler() (rh *RequestHandler) {
	l, _ := zap.NewProduction()

	cfg, err := parseConfig("")
	if err != nil {
		l.Error("failed to parse config", zap.Error(err))
		os.Exit(1)
	}

	pgConn, err := pgx.Connect(context.Background(), getPgUrlByConfig(cfg))
	if err != nil {
		l.Error("failed to connect to pg", zap.Error(err))
		os.Exit(1)
	}

	cache := redis.NewClient(&redis.Options{
		Addr:     cfg.GetString("redis.addr"),
		Password: cfg.GetString("redis.password"),
		DB:       cfg.GetInt("redis.db"),
	})

	pairsRepo := entities.NewCryptoInfoRepoCache(pgConn, cache)

	rh = NewRequestHandler(l, cfg, pairsRepo, cache)

	return rh
}

func (rh *RequestHandler) priceHandler(w http.ResponseWriter, r *http.Request) {
	fsyms := r.FormValue("fsyms")
	tsyms := r.FormValue("tsyms")

	var (
		err  error
		resp []byte
		ctx  = context.Background()
	)

	if resp, err = rh.cryptoInfo.GetLatestCryptoInfo(ctx); err != nil {
		resp, err = rh.getCryptoInfo(ctx, fsyms, tsyms)
		if err != nil {
			rh.l.Error("getCryptoInfo error", zap.Error(err))
			return
		}
	}

	var rw entities.CryptocompareResp
	if err = json.Unmarshal(resp, &rw); err != nil {
		rh.l.Error("failed to unmarshal json", zap.Error(err))
		http.Error(w, "Something bad happened", http.StatusInternalServerError)
		return
	}

	clearResp := rw.GetByFsymsAndTsyms(fsyms, tsyms)

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(clearResp); err != nil {
		rh.l.Error("w.Write error", zap.Error(err))
	}

	return
}
