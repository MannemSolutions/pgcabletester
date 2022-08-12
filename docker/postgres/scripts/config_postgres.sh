#!/bin/bash
function print_config() {
PGBIN=/usr/pgsql-14/bin
echo "
# Generic
# STKEEPER
PGDATA=/pgdata/data
PGCERTS=/pgdata/certs
PGPORT=5432
PGBIN=$PGBIN
PATH=/host/bin:$PGBIN:~/.local/bin:$PATH"
}

#MYIP=$(ip a | grep -oE 'inet ([0-9]{1,3}\.){3}[0-9]{1,3}' | sed -e '/127\.0\.0\.1/d' -e 's/inet //')
#MYHOSTNAME=$(host "${MYIP}" | sed -e 's/.* //' -e 's/\..*//')

print_config
