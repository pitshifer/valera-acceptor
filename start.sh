#!/bin/bash

set -e 

sed -i "s|{{POSTGRES_PASSWORD}}|$POSTGRES_PASSWORD|" /app/configs/apiserver.toml

exec "$@"
