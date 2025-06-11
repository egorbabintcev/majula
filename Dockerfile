FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o /majula cmd/majula/main.go

FROM alpine:3
COPY --from=builder majula /bin/majula
ENTRYPOINT ["/bin/majula"]