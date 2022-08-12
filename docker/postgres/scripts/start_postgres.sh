#!/bin/bash

function pgconfig() {
echo "
listen_addresses='*'
ssl = on
ssl_cert_file = '/pgdata/certs/tls/int_server/certs/dbdiff_postgres_1.pem'
ssl_key_file = '/pgdata/certs/tls/int_server/private/dbdiff_postgres_1.key.pem'
ssl_ca_file = '/pgdata/certs/tls/int_client/certs/ca-chain-bundle.cert.pem'
"
}


SCRIPTDIR=$(dirname $0)
eval $($SCRIPTDIR/config_postgres.sh | sed -e 's/#.*//' -e '/[a-zA-Z0-9]/!d' -e 's/^/export /')
if [ "$(id -un)" != postgres ]; then
  mkdir -p "${PGDATA}" "${PGCERTS}"
  chown postgres: "${PGDATA}" "${PGCERTS}"
  su - postgres $0
  exit
fi
pip3 install --user chainsmith
chainsmith -c /host/chainsmith.yml
cp /pgdata/certs/tls/int_client/certs/postgres.pem /host/certs/postgres.crt
cp /pgdata/certs/tls/int_client/private/postgres.key.pem /host/certs/postgres.key

initdb --auth-host=cert
sed -i 's/^host /hostssl/' /pgdata/data/pg_hba.conf
echo 'hostssl   all             all             0.0.0.0/0            cert' >> /pgdata/data/pg_hba.conf
pgconfig >> "$PGDATA/postgresql.conf"
postgres
while /bin/true; do
  sleep 10
done
