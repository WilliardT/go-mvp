FROM golang:1.25.6-bookworm AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/cmd/go-mvp-app/exe /app/cmd/go-mvp-app

FROM alpine:3.23
WORKDIR /app
COPY --from=builder /app/cmd/go-mvp-app/exe /app
CMD ["/app/exe"]
