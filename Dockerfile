FROM golang:alpine AS builder
ENV GOPROXY https://goproxy.cn, direct
WORKDIR /home
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o main .
FROM alpine:latest
COPY --from=builder /home/main /home/main
CMD ["/home/main"]