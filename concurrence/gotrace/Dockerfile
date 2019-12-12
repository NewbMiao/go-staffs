FROM divan/golang:gotrace
ENV GOOS=linux
RUN go get -u github.com/divan/gotrace && cd $GOPATH/src/github.com/divan/gotrace && git checkout go18 && go install
RUN mv /go/bin/gotrace /usr/local/bin
ADD https://gist.githubusercontent.com/NewbMiao/f4340b483e3dfc057911cba8e7a37562/raw/a009b89413b5ed1052a4108916c7c042eb0808ce/gotrace.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
EXPOSE 2000
ENTRYPOINT ["/entrypoint.sh"]