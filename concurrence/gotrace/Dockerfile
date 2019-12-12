FROM divan/golang:gotrace
RUN go get -u github.com/divan/gotrace && cd $GOPATH/src/github.com/divan/gotrace && git checkout go18 && go install
RUN mv /go/bin/gotrace /usr/local/bin
EXPOSE 2000