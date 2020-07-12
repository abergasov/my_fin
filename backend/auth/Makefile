COMMIT?=$(shell git ls-files | xargs sha256sum | cut -d" " -f1 | sha256sum | cut -d" " -f1)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

build:
	@echo "-- building auth binary"
	go build -ldflags "-X main.buildHash=${COMMIT} -X main.buildTime=${BUILD_TIME}" -o ./bin/my_fin_auth