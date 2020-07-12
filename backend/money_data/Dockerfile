# 1 шаг - сборка пакета
FROM golang:1.14 AS build
COPY . /go/src/my_fin
WORKDIR /go/src/my_fin/backend/money_data
RUN make build

# 2 шаг - создание контейнера для запуска
FROM debian:buster-slim
WORKDIR /app
#COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/src/my_fin/backend/money_data/bin/* /app/
COPY --from=build /go/src/my_fin/backend/money_data/config/.env /app/config/.env
COPY --from=build /go/src/my_fin/scripts/wait-for-it.sh /app/

RUN chmod +x /app/*
EXPOSE 8080/tcp
CMD /app/my_fin