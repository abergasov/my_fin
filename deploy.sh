#!/bin/bash

git fetch --all
git reset --hard origin/master
go mod vendor

echo "migrate database"

cd helpers || exit
composer install
if php app.php migrate ; then
    echo "migrated"
else
    echo "migrated failed"
    exit 1
fi

cd ..

echo "build binary"

cd backend || exit

if make build; then
  echo "build ok"
  sudo service my_fin restart
else
  echo "build failed"
  exit 1
fi

cd ..

echo "build front js"

cd front || exit
npm ci
npm run build