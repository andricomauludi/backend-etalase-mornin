FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=auto
RUN cd /build && git clone https://github.com/andricomauludi/backend-etalase-mornin.git

RUN cd /build/backend-etalase-mornin && go build

EXPOSE 8090

ENTRYPOINT ["/build/backend-etalase-mornin/main"]