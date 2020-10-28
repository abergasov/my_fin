#!/usr/bin/env bash

clean() {
  echo "stop containers";
  docker container stop app_main.mf webserver.mf dbMysql.mf
  echo "drop containers"
  docker rm -v app_main.mf webserver.mf dbMysql.mf
}

clean

git fetch --all
git reset --hard origin/master

echo "migrate database"

cd helpers || exit
composer install
if php app.php migrate ; then
    echo "migrated"
else
    echo "migrated failed"
    exit 1
fi

echo "run prod config"
echo "RUN docker-compose.yml "
docker-compose -f docker-compose.yml pull
docker-compose -f docker-compose.yml up --build
