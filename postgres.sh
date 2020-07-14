#!/bin/bash

db=$1
user=$2
pass=$3

docker run --name some-postgres \
  -e POSTGRES_DB="$db" \
  -e POSTGRES_USER="$user" \
  -e POSTGRES_PASSWORD="$pass" \
  -d postgres