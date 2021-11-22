package entities

import (
	"encoding/json"
	"strings"
)

type fCurrency string
type tCurrency string

type CryptocompareResp struct {
	Raw     map[fCurrency]tCurrencyInfo        `json:"RAW,omitempty"`
	Display map[fCurrency]tCurrencyInfoDisplay `json:"DISPLAY,omitempty"`
}

type tCurrencyInfo map[tCurrency]CurrencyInfo
type tCurrencyInfoDisplay map[tCurrency]CurrencyInfoDisplay

// GetByFsymsAndTsyms get only needed fsyms and tsyms
func (cc *CryptocompareResp) GetByFsymsAndTsyms(fsymsRaw, tsymsRaw string) []byte {
	fsyms := strings.Split(fsymsRaw, ",")
	tsyms := strings.Split(tsymsRaw, ",")

	resp := CryptocompareResp{
		Raw:     cc.getByFsymsAndTsyms(fsyms, tsyms),
		Display: cc.getByFsymsAndTsymsDisplay(fsyms, tsyms),
	}

	respStr, _ := json.Marshal(resp)

	return respStr
}

func (cc *CryptocompareResp) getByFsymsAndTsyms(
	fsyms,
	tsyms []string,
) (raw map[fCurrency]tCurrencyInfo) {
	raw = make(map[fCurrency]tCurrencyInfo)

	var (
		fsymOk, tsymOk bool
		currencyInfo   CurrencyInfo
	)

	for _, fsym := range fsyms {
		fCurr := fCurrency(fsym)
		if _, ok := raw[fCurr]; !ok {
			raw[fCurr] = make(map[tCurrency]CurrencyInfo)
		}

		for _, tsym := range tsyms {
			tCurr := tCurrency(tsym)
			if _, fsymOk = cc.Raw[fCurr]; fsymOk {
				if currencyInfo, tsymOk = cc.Raw[fCurr][tCurr]; tsymOk {
					raw[fCurr][tCurr] = currencyInfo
				}
			}
		}
	}

	return raw
}

func (cc *CryptocompareResp) getByFsymsAndTsymsDisplay(
	fsyms,
	tsyms []string,
) (display map[fCurrency]tCurrencyInfoDisplay) {
	display = make(map[fCurrency]tCurrencyInfoDisplay)

	var (
		fsymOk, tsymOk      bool
		currencyInfoDisplay CurrencyInfoDisplay
	)

	for _, fsym := range fsyms {
		fCurr := fCurrency(fsym)

		if _, ok := display[fCurr]; !ok {
			display[fCurr] = make(map[tCurrency]CurrencyInfoDisplay)
		}
		for _, tsym := range tsyms {
			tCurr := tCurrency(tsym)
			if _, fsymOk = cc.Display[fCurr][tCurr]; fsymOk {
				if currencyInfoDisplay, tsymOk = cc.Display[fCurr][tCurr]; tsymOk {
					display[fCurr][tCurr] = currencyInfoDisplay
				}
			}
		}
	}

	return display
}

type CurrencyInfo struct {
	CHANGE24HOUR    float64 `json:"CHANGE24HOUR"`
	CHANGEPCT24HOUR float64 `json:"CHANGEPCT24HOUR"`
	OPEN24HOUR      float64 `json:"OPEN24HOUR"`
	VOLUME24HOUR    float64 `json:"VOLUME24HOUR"`
	VOLUME24HOURTO  float64 `json:"VOLUME24HOURTO"`
	LOW24HOUR       float64 `json:"LOW24HOUR"`
	HIGH24HOUR      float64 `json:"HIGH24HOUR"`
	PRICE           float64 `json:"PRICE"`
	SUPPLY          float64 `json:"SUPPLY"`
	MKTCAP          float64 `json:"MKTCAP"`
}

type CurrencyInfoDisplay struct {
	CHANGE24HOUR    string `json:"CHANGE24HOUR"`
	CHANGEPCT24HOUR string `json:"CHANGEPCT24HOUR"`
	OPEN24HOUR      string `json:"OPEN24HOUR"`
	VOLUME24HOUR    string `json:"VOLUME24HOUR"`
	VOLUME24HOURTO  string `json:"VOLUME24HOURTO"`
	LOW24HOUR       string `json:"LOW24HOUR"`
	HIGH24HOUR      string `json:"HIGH24HOUR"`
	PRICE           string `json:"PRICE"`
	SUPPLY          string `json:"SUPPLY"`
	MKTCAP          string `json:"MKTCAP"`
}
