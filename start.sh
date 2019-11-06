#!/bin/bash

set -e 

sed -i "s|{{POSTGRES_PASSWORD}}|$POSTGRES_PASSWORD|" /app/configs/acceptor.toml

/app/acceptor -migrate=up 
exec "$@"
