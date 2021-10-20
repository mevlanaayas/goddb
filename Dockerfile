FROM golang:1.16.0-alpine AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get -v -t ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder app .
EXPOSE 8099
CMD ["./app"]
