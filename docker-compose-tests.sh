#!/bin/bash
set -e

docker-compose down --remove-orphans || echo new or partial install
docker-compose up -d
