# First stage: build the Go application
FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Second stage: start from scratch
FROM scratch
COPY --from=builder /app/main .
ENTRYPOINT ["./main"]
CMD ["-help"]