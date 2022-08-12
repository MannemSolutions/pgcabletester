#!/bin/bash
SCRIPTDIR=$(dirname $0)
for SVC in postgres; do
  "${SCRIPTDIR}/start_${SVC}.sh" &
  sleep 1
done
wait
