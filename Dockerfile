FROM alpine:latest AS certificates_builder
RUN apk --no-cache add tzdata ca-certificates  


FROM golang:1.21.5-alpine3.18 AS builder
ADD . /go/stress-test

WORKDIR /go/stress-test

RUN mkdir deploy
RUN go clean --modcache
RUN go mod tidy

RUN CGO_ENABLED=0 go build -o stress-test main.go 
RUN mv stress-test ./deploy/stress-test

FROM scratch AS production

COPY --from=certificates_builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=certificates_builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /go/stress-test/deploy /stress-test/

WORKDIR /stress-test

ENTRYPOINT  ["./stress-test"]