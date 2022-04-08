FROM golang:1.18.0-alpine3.15 AS build
# Support CGO and SSL
RUN apk add --update sudo
RUN apk --no-cache add gcc g++ make
RUN apk add git
WORKDIR /go/src/app
COPY . .
RUN go env -w GO111MODULE=off
RUN go get github.com/gorilla/mux
RUN go get github.com/go-sql-driver/mysql
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/test ./main.go
# Setting web service
FROM alpine:3.15
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/app/bin /go/bin
EXPOSE 8080
ENTRYPOINT /go/bin/test --port 8080