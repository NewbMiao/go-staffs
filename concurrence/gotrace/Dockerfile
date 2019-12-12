FROM divan/golang:gotrace
ENV GOOS=linux
RUN go get -u github.com/divan/gotrace && cd $GOPATH/src/github.com/divan/gotrace && git checkout go18 && go install
RUN mv /go/bin/gotrace /usr/local/bin
ADD https://gist.githubusercontent.com/NewbMiao/f4340b483e3dfc057911cba8e7a37562/raw/8a71c3134a3bd048af93a70ff1aa66f909e9e1b7/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
EXPOSE 2000
ENTRYPOINT ["/entrypoint.sh"]