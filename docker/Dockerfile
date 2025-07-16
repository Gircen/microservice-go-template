FROM alpine:3.22.0
LABEL authors="antonkolesnikov"
RUN apk add --no-cache git make musl-dev go

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

RUN mkdir /etc/ms/
RUN chmod -R 777 /etc/ms/

CMD ["git","clone","https://github.com/Gircen/microservice-go-template","/etc/ms"]
RUN chmod -R 777 /etc/ms/
CMD ["chmod", "+", "/etc/ms/microservice-go-template/run.sh"]
RUN /etc/ms/microservice-go-template/run.sh
