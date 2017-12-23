FROM golang:1.8
COPY . "$GOPATH/src/github.com/YlingMA/Agenda-Service2"
RUN cd "$GOPATH/src/github.com/YlingMA/Agenda-Service2/cli" && go get -v && go install -v
RUN cd "$GOPATH/src/github.com/YlingMA/Agenda-Service2/service" && go get -v && go install -v
WORKDIR /
EXPOSE 8080
VOLUME [ "/data" ]
