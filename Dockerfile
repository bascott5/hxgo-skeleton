# Build
FROM golang:1.23.3-alpine AS build-env
ENV APP_NAME=hxgo-skeleton
ENV CMD_PATH=app.go
COPY . $GOPATH/src/$APP_NAME
COPY .env $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME

# Building the Go application
RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH
 
# Run
FROM alpine:3.14
ENV APP_NAME=hxgo-skeleton
COPY --from=build-env /$APP_NAME .
COPY .env .
EXPOSE 8080
CMD ./$APP_NAME
