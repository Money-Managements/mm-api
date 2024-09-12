#!/bin/bash

if [ "$1" == "env" ]; then
  source ./script/env.sh
elif [ "$1" == "run:dev" ]; then
  source ./script/env.sh
  go run ./app/main.go
else
  echo "Usage: project [comand]"
fi