FROM golang

ENV GOPATH /opt/go:$GOPATH
ENV PATH /opt/go/bin:$PATH
ADD . /opt/go/src/github.com/godocapi
WORKDIR /opt/go/src/github.com/godocapi

RUN go get github.com/derekparker/delve/cmd/dlv
RUN go build -o main main.go
#RUN dlv debug github.com/godocapi -l 0.0.0.0:2345 --headless=true --log=true -- server
CMD ["./main"]
