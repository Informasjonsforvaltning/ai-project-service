FROM golang:1.24-alpine AS build-env

ENV APP_NAME=ai-project-service
ENV CMD_PATH=main.go

COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME

RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH

FROM alpine:3.21

ENV APP_NAME=ai-project-service
ENV GIN_MODE=release

COPY --from=build-env /$APP_NAME .
COPY *.csv .

EXPOSE 8080

CMD ["/ai-project-service"]
