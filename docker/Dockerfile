FROM alpine:3.22.0
LABEL authors="antonkolesnikov"
RUN apk add --no-cache git make musl-dev go

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH
ENV MSROOT /tmp/mc

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

RUN mkdir /opt/ms/
RUN chmod -R 777 /opt/ms/
RUN mkdir /tmp/ms/


RUN git clone https://github.com/Gircen/microservice-go-template tmp/ms
RUN chmod + ./tmp/ms/sh/prod_build.sh
RUN chmod -R 777 /tmp/ms/
RUN cp ./tmp/ms/sh/prod_build.sh /tmp/ms/
RUN sh ./tmp/ms/prod_build.sh
RUN chmod -R 777 /tmp/ms/build/

RUN mv tmp/ms/build/* opt/ms

CMD ["chmod", "+", "/opt/ms/microservice-go-template/run.sh"]
RUN /opt/ms/microservice-go-template/sh/build.sh
