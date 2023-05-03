# MAINTAINER Jonathan Mainguy <jon@soh.re>
FROM golang:1.20.4
WORKDIR /go/src/app
ENV GO111MODULE=on
ADD . .
RUN go build
EXPOSE 8080
CMD ["/go/src/app/httpIdle"]
