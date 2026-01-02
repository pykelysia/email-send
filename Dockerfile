FROM golang:1.24-alpine AS builder
WORKDIR /src

RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /emailsend ./emailsend.go

# Runtime stage
FROM alpine:3.18
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /emailsend /app/emailsend

# Copy configuration and other runtime assets that the binary may need
COPY develop.yaml ./

USER 65532:65532
EXPOSE 8080
ENTRYPOINT ["/app/emailsend"]
