#!/bin/bash

git fetch --all
git reset --hard origin/master
go mod vendor

cd helpers || exit
composer install
if php app.php migrate ; then
    echo "migrated"
else
    echo "migrated failed"
    exit 1
fi

cd ..

if go build main.go ; then
  echo "build ok"
  sudo service my_fin restart
else
  echo "build failed"
  exit 1
fi

cd front || exit
npm ci
npm run build