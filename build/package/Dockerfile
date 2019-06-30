FROM golang:1.11.10-alpine3.9
LABEL maintainer="Siddhartha Basu <siddhartha-basu@northwestern.edu>"
RUN apk add --no-cache git build-base
RUN mkdir -p /slack-notify
WORKDIR /slack-notify
ADD cmd cmd
ADD internal internal
COPY go.mod go.sum ./
RUN go get ./... && \
    go build -o app cmd/slack-notify/main.go

FROM alpine:3.9
RUN apk --no-cache add ca-certificates
COPY --from=0 /slack-notify/app /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/app"]
