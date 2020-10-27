FROM golang:1.14 AS build
COPY . /go/src/my_fin
WORKDIR /go/src/my_fin
RUN make build

# 2 шаг - создание контейнера для запуска
FROM debian:buster-slim
WORKDIR /app
COPY --from=build /go/src/my_fin/bin/* /app/
COPY --from=build /go/src/my_fin/configs/* /app/configs/
COPY --from=build /go/src/my_fin/scripts/* /app/configs/
RUN chmod +x /app/*
EXPOSE 29010/tcp
CMD /app/archivarius