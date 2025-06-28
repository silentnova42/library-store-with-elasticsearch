FROM golang:1.24.4 AS builder
WORKDIR /root

COPY go.mod .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOSE=linux go build -o main ./cmd/main.go

FROM alpine:3.22.0
WORKDIR /root

COPY --from=builder /root/main .
CMD [ "./main" ]
