FROM golang:1.24.0-alpine 

RUN apk update && apk upgrade && \
	apk add bash ca-certificates

COPY . /go/src/github.com/ardanlabs/tour/

WORKDIR /go/src/github.com/ardanlabs/tour

RUN go build -o /opt/tour/tour ./cmd/tour

WORKDIR /opt/tour

EXPOSE 8080

CMD ["./tour", "-http", "0.0.0.0:8080", "-origin", "tour.ardanlabs.com:443", "-scheme", "https", "-openbrowser", "false"]
