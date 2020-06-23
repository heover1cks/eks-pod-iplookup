FROM golang:1.14 AS build
COPY . /app
WORKDIR /app
RUN go get -d github.com/gorilla/mux/
RUN go build -o app
FROM alpine:latest AS runtime
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
WORKDIR /app
COPY --from=build /app .
CMD ["./app"]
