package repository

import (
	"context"

	"github.com/jackc/pgx/v4"
)

var sqlAddCryptoInfo = `
INSERT INTO cryptocompare_pairs(raw) VALUES($1)
`

func (mr *CryptoInfoRepo) AddCryptoInfo(ctx context.Context, cryptoInfo string) (err error) {
	_, err = mr.PgConn.Exec(ctx, sqlAddCryptoInfo, cryptoInfo)
	return err
}

var sqlGetLatestCryptoInfo = `
SELECT raw FROM cryptocompare_pairs ORDER BY id DESC LIMIT 1
`

func (mr *CryptoInfoRepo) GetLatestCryptoInfo(ctx context.Context) (latestInfo []byte, err error) {
	var cryptoInfo string
	if err = mr.PgConn.QueryRow(ctx, sqlGetLatestCryptoInfo).Scan(&sqlGetLatestCryptoInfo); err != nil {
		return nil, err
	}

	return []byte(cryptoInfo), nil
}

type CryptoInfoRepo struct {
	PgConn *pgx.Conn
}
