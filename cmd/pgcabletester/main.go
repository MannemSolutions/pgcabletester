package main

import (
	"github.com/mannemsolutions/pgcabletester/internal"
	"github.com/mannemsolutions/pgcabletester/pkg/cbl"
)

func main() {
	var err error
	var config internal.PCTConfig
	initLogger()
	if config, err = internal.NewConfig(); err != nil {
		log.Fatal(err)
	} else {
		cbl.InitLogger(log)
		enableDebug(config.Debug)
		defer log.Sync() //nolint:errcheck
	}
	if config.ServerMode {
		cbl.RunServers(config.Hosts)
	} else {
		cbl.RunClients(config)
	}
}
