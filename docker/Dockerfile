FROM alpine:3.22.0
LABEL authors="antonkolesnikov"
RUN apk add --no-cache git make musl-dev go

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

RUN mkdir /opt/ms/
RUN chmod -R 777 /opt/ms/

CMD ["git","clone","https://github.com/Gircen/microservice-go-template","/opt/ms"]

CMD ["chmod", "+", "/opt/ms/microservice-go-template/run.sh"]
RUN /opt/ms/microservice-go-template/sh/build.sh
