FROM golang:1.13 as builder

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -v /app/cmd/acceptor

#################

FROM alpine:latest

RUN apk --no-cache add ca-certificates
RUN apk --no-cache add bash

RUN mkdir /app
COPY --from=builder /app /app
WORKDIR /app

EXPOSE 8080

ENTRYPOINT [ "/app/start.sh" ]

CMD ["/app/acceptor"]
