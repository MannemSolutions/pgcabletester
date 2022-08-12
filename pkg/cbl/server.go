package cbl

import (
	"context"
	"fmt"
	wire "github.com/jeroenrinzema/psql-wire"
	"github.com/mannemsolutions/pgcabletester/internal"
	"net"
)

func RunServers(hosts internal.PCTConfigHosts) {
	for hostname, config := range hosts {
		if ips, err := net.LookupIP(hostname); err != nil {
			log.Fatalf("Could not get ip for %s, %e", hostname, err)
		} else {
			for _, ip := range ips {
				for _, port := range config.Ports {
					address := fmt.Sprintf("%s:%d", ip.String(), port)
					go wire.ListenAndServe(address, func(ctx context.Context, query string, writer wire.DataWriter) error {
						fmt.Println(query)
						return writer.Complete("OK")
					})
				}
			}
		}
	}
}
