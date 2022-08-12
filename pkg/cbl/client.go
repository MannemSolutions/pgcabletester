package cbl

import (
	"fmt"
	"github.com/mannemsolutions/pgtester/pkg/pg"
	"go.uber.org/zap"
	"net"
)

func InitPgLogger(logger *zap.SugaredLogger) {
	pg.Initialize(logger)
}

type CBLClient struct {
	Dsn    pg.Dsn
	tables []string
}

type CBLClients []CBLClient

func (cblCs CBLClients) Hosts() (hosts []string) {
	// using a map for easier checking if it was already in there
	_hosts := make(map[string]bool)
	for _, cblc := range cblCs {
		if hostName, exists := cblc.Dsn["host"]; !exists {
			continue
		} else if _, exists = _hosts[hostName]; exists {
			continue
		} else {
			_hosts[hostName] = true
		}
	}
	for hostName := range _hosts {
		hosts = append(hosts, hostName)
	}
	return hosts
}

/*
- resolve dns
- connect to port
- number of hops
- open connection
- log in
- check table
*/

func RunClients(cblCs CBLClients) {
	errors := 0
	for _, hostName := range cblCs.Hosts() {
		log.Infof("resolving in DNS")
		if _, err := net.LookupIP(hostName); err != nil {
			log.Errorf("could not resolve: %s", err.Error())
			errors += 1
		}
	}
	if errors > 0 {
		log.Fatal("Too many errors!!!")
	}
	for _, cblc := range cblCs {
		conn := pg.NewConn(cblc.Dsn, 0, 0)
		for _, table := range cblc.tables {
			conn.GetOneField(fmt.Sprintf("select * from "))
		}
	}

}
