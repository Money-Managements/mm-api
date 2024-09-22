#!/bin/bash

set_envs() {
  export DB_USER="user"
  export DB_PASSWORD="root"
  export DB_HOST="money-manager-db-1"
  export DB_PORT="3306"
  export DB_NAME="money_manager"
}

start_dev() {
  set_envs
  go run ./app/main.go
}

seed_dev() {
  set_envs
  go run ./script/seed/main.go
}

if [ "$1" == "start:dev" ]; then
  start_dev
elif [ "$1" == "set:envs" ]; then
  set_envs
elif [ "$1" == "seed:dev" ]; then
  seed_dev
else
  echo "Usage [command]"
fi