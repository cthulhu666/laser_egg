FROM golang:1.15.7 as build
ENV HOME /opt/app
COPY . $HOME
WORKDIR $HOME

RUN go build cmd/main.go && cp ./main /go/bin/

FROM debian:buster
ENV HOME /opt/app
WORKDIR $HOME

COPY --from=build /go/bin/ /go/bin/

ENTRYPOINT ["/go/bin/main"]
