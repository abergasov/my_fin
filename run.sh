#!/usr/bin/env bash

clean() {
  echo "stop containers";
  docker stop app_main.mf app_auth.mf webserver.mf dbMysql.mf
  echo "drop containers"
  docker rm -v app_main.mf app_auth.mf webserver.mf dbMysql.mf
}

clean

docker-compose -f docker-compose.yml pull

docker-compose -f docker-compose.yml up --build
