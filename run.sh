#!/usr/bin/env bash

env=${env:-dev}
s=${s:-""}

while [ $# -gt 0 ]; do

   if [[ $1 == *"--"* ]]; then
        param="${1/--/}"
        declare "$param"="$2"
        # echo $1 $2 // Optional to see the parameter:value result
   fi

  shift
done

echo "${s}"
echo "Enviroment: ${env}"

clean() {
  echo "stop containers";
  docker container stop app_main.mf webserver.mf dbMysql.mf
  echo "drop containers"
  docker rm -v app_main.mf webserver.mf dbMysql.mf
}

clean

if [ "$env" == "dev" ]; then
  serviceList="webserver dbMysql $s"
  echo "RUNNING SERVICES: $serviceList"
  echo "RUN docker-compose-dev.yml "
  docker-compose -f docker-compose-dev.yml pull
  docker-compose -f docker-compose-dev.yml up --build ${serviceList}
else
  echo "run prod config"
  echo "RUN docker-compose.yml "
  docker-compose -f docker-compose.yml pull
  docker-compose -f docker-compose.yml up --build
fi