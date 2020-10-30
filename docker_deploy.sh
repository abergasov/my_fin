#!/usr/bin/env bash

clean() {
  echo "stop containers";
  docker container stop app_main.mf dbMysql.mf
  echo "drop containers"
  docker rm -v app_main.mf dbMysql.mf
  echo "drop old images"
  docker rmi $(docker images -f dangling=true -q)
}

clean

echo "run prod config"
echo "RUN docker-compose.yml "
docker-compose -f docker-compose.yml pull
docker-compose -f docker-compose.yml up --build
