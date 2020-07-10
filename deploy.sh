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

BUILD_HASH=$(git ls-files | xargs sha256sum | cut -d" " -f1 | sha256sum | cut -d" " -f1)
BUILD_TIME=$(date -u '+%Y-%m-%d_%H:%M:%S')
echo "$BUILD_TIME"
echo "$BUILD_HASH"

if go build -ldflags="-X 'main.buildTime=${BUILD_TIME}' -X 'main.buildHash=${BUILD_HASH}'"; then
  echo "build ok"
  sudo service my_fin restart
else
  echo "build failed"
  exit 1
fi

cd front || exit
npm ci
npm run build