FROM golang:1.23.4-alpine as build

COPY . /app

WORKDIR /app

RUN go mod init go-web-server

RUN go mod tidy

RUN go build -o main main.go

FROM scratch

COPY --from=build /app/main /app/main

COPY --from=build /app/index.html /app/index.html

WORKDIR /app

CMD ["./main"]