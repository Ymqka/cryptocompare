package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

const cryptoCompareUrl = "https://min-api.cryptocompare.com/data/pricemultifull?fsyms=%s&tsyms=%s"

// getCryptoInfo get price of crypt in raw json
// from min-api.cryptocompare.com/data/pricemultifull
func (rh *RequestHandler) getCryptoInfo(
	ctx context.Context,
	fsyms, tsyms string,
) (cryptoInfo []byte, err error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf(cryptoCompareUrl, fsyms, tsyms),
		nil,
	)
	if err != nil {
		return []byte{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}
