package main

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	defaultConfigPath = "cfg/config.yaml"
	tsymsCfgKey       = "tsyms"
	fsymsCfgKey       = "fsyms"
)

func parseConfig(path string) (cfg *viper.Viper, err error) {
	cfg = viper.New()

	if path == "" {
		path = defaultConfigPath
	}
	cfg.SetConfigFile(path)
	cfg.SetConfigType("yaml")

	if err = cfg.ReadInConfig(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (rh *RequestHandler) getTsymsFromConfig() (tsyms string) {
	cfgTsyms := rh.cfg.GetStringSlice(tsymsCfgKey)
	for i, t := range cfgTsyms {
		tsyms += t
		if i+1 != len(cfgTsyms) {
			tsyms += ","
		}
	}

	return tsyms
}

func (rh *RequestHandler) getFsymsFromConfig() (fsyms string) {
	cfgFsyms := rh.cfg.GetStringSlice(fsymsCfgKey)
	for i, t := range cfgFsyms {
		fsyms += t
		if i+1 != len(cfgFsyms) {
			fsyms += ","
		}
	}

	return fsyms
}

func getPgUrlByConfig(cfg *viper.Viper) (url string) {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.GetString("psql.login"),
		cfg.GetString("psql.password"),
		cfg.GetString("psql.hostname"),
		cfg.GetInt("psql.port"),
		cfg.GetString("psql.db"),
	)
}
