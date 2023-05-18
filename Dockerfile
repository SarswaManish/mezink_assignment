FROM golang:1.20
RUN mkdir /assignment
WORKDIR /assignment
COPY . /assignment

RUN go build  -o ./bin/server ./cmd/server

CMD ["/assignment/bin/server", "--config-file", "/assignment/configs/dev.toml", "--debug=false"]
