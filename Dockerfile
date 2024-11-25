# Build
FROM golang:1.23.3-alpine as build-env
ENV APP_NAME hxgo-skeleton
ENV CMD_PATH app.go
COPY . $GOPATH/src/$APP_NAME
COPY .env $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME

# Downloading and building Tailwind
RUN <<EOF
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.1/tailwindcss-linux-x64
chmod +x tailwindcss-linux-x64
mv tailwindcss-linux-x64 tailwindcss
tailwindcss -i $GOPATH/src/$APP_NAME/web/styles/input.css -o $GOPATH/src/$APP_NAME/web/styles/output.css
EOF

# Building the Go application
RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH
 
# Run
FROM alpine:3.14
ENV APP_NAME hxgo-skeleton
COPY --from=build-env /$APP_NAME .
COPY .env .
EXPOSE 42069
CMD ./$APP_NAME
