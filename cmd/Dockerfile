FROM golang:1.15 AS build
COPY . /go/src/my_fin
WORKDIR /go/src/my_fin
RUN make build

# 2 шаг - создание контейнера для запуска
FROM debian:buster-slim
WORKDIR /app
COPY --from=build /go/src/my_fin/bin/* /app/
COPY --from=build /go/src/my_fin/configs/* /app/configs/
COPY --from=build /go/src/my_fin/scripts/wait-for-it.sh /app/wait-for-it.sh
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
RUN chmod +x /app/*
EXPOSE 8080/tcp
CMD /app/my_fin