FROM golang:1.11.1-alpine3.8
ARG PROJECT_DIR
WORKDIR /go/src/${PROJECT_DIR}
ADD . /go/src/${PROJECT_DIR}
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ./server ./cmd/server

FROM alpine:3.8
ARG PROJECT_DIR
ARG DEFAULT_PORT

WORKDIR /
RUN apk add --no-cache ca-certificates
RUN apk add --update openssl
# RUN openssl genrsa -out rsa_key 2048
# RUN openssl rsa -in rsa_key -pubout > rsa_key.pub

COPY --from=0 /go/src/${PROJECT_DIR}/server /server
COPY --from=0 /go/src/${PROJECT_DIR}/docker-entrypoint.sh /
EXPOSE ${DEFAULT_PORT}
ENTRYPOINT ["sh", "/docker-entrypoint.sh"]