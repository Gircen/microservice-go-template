FROM alpine:3.22.0
LABEL authors="antonkolesnikov"
RUN apk add --no-cache git make musl-dev go

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH
ENV MSROOTTMP /tmp/microservice-go-template

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

RUN mkdir ${MSROOTTMP}


RUN git clone https://github.com/Gircen/microservice-go-template ${MSROOTTMP}
RUN chmod + ${MSROOTTMP}/sh/prod_build.sh
RUN chmod -R 777 ${MSROOTTMP}

RUN cp ${MSROOTTMP}/sh/prod_build.sh ${MSROOTTMP}
RUN sh ${MSROOTTMP}/prod_build.sh
RUN chmod -R 777 ${MSROOTTMP}/build/

RUN mkdir /opt/microservice-go-template/
RUN chmod -R 777 /opt/ms/
RUN mv tmp/ms/build/* opt/ms
RUN rm -rf tmp/ms/*

RUN /opt/ms/microservice-go-template
