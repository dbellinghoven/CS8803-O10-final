FROM golang:1.14.6-alpine3.12

WORKDIR /go/src/app

COPY cleanup.go cleanup.sh ./

RUN chmod +x ./cleanup.sh

CMD ["./cleanup.sh"]
