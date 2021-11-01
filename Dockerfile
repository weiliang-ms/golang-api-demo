FROM golang:1.16.0-alpine AS builder
WORKDIR /work
ADD . .
RUN go build

FROM alpine
WORKDIR /work
COPY --from=builder /work/_output/golang-api-demo /api-server
CMD ["/api-server"]