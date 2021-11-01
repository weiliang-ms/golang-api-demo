FROM golang:1.16.0-alpine AS builder
WORKDIR /work
ADD . .
RUN go build -o _output/api-server

FROM alpine
WORKDIR /work
COPY --from=builder /work/_output/api-server /api-server
CMD ["/api-server"]