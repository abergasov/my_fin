# 1 шаг - сборка пакета
FROM golang:1.14 AS build
COPY . /go/src/my_fin
WORKDIR /go/src/my_fin/backend/auth
RUN make build

# 2 шаг - создание контейнера для запуска
FROM debian:buster-slim
WORKDIR /app
#COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/src/my_fin/backend/auth/bin/* /app/
COPY --from=build /go/src/my_fin/backend/auth/config/.env /app/config/.env
COPY --from=build /go/src/my_fin/scripts/wait-for-it.sh /app/

RUN chmod +x /app/*
EXPOSE 8090/tcp
CMD /app/my_fin_auth