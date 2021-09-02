FROM golang:1.16-alpine AS builder
RUN apk update && apk add --no-cache make git
WORKDIR /go/src/github.com/bro-n-bro/osjuno
COPY . ./
RUN make build

FROM alpine:latest
COPY --from=builder /go/src/github.com/bro-n-bro/osjuno/build/osjuno /usr/bin/osjuno
RUN osjuno init --home /osjuno
# todo entrypoint
CMD [ "osjuno", "parse", "--home", "/osjuno"]