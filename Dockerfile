FROM golang:1.16-alpine AS builder
RUN apk update && apk add --no-cache make git
WORKDIR /go/src/github.com/bro-n-bro/osjuno
COPY . ./
RUN make build

FROM alpine:latest
WORKDIR osjuno
COPY --from=builder /go/src/github.com/bro-n-bro/osjuno/build/osjuno /usr/bin/osjuno
COPY ./entrypoint.sh ./
ENTRYPOINT [ "./entrypoint.sh" ]
CMD [ "osjuno" ]
