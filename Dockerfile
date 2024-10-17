FROM golang:1.23.2-bookworm AS builder
WORKDIR /app
ADD . /app
RUN go build -o bin/server server/main.go


FROM debian:bookworm-slim
COPY --from=builder /app/bin/server .
EXPOSE 8080
CMD ["./server"]
