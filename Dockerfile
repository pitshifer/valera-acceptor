FROM golang:1.13 as builder

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -v /app/cmd/apiserver

#################

FROM alpine:latest

RUN apk --no-cache add ca-certificates

# WORKDIR /root/

RUN mkdir /app
COPY --from=builder /app /app
WORKDIR /app

EXPOSE 8080

CMD ["/app/apiserver"]
