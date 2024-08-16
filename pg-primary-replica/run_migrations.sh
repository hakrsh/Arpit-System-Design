#!/bin/bash

if [ -f .env ]; then
  eval $(grep -v '^#' .env | xargs)
fi

migrate -path db/migrations -database $POSTGRES_PRIMARY_SOURCE -verbose up