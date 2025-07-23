FROM alpine:3.22.0
LABEL authors="antonkolesnikov"
RUN apk add --no-cache git make musl-dev go

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH
ENV MSROOTTMP /tmp/microservice-go-template
ENV MSROOTOPT /opt/ms
ENV MS_CONFIG /opt/ms/config/configmap.yaml

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

RUN mkdir ${MSROOTTMP}


RUN git clone https://github.com/Gircen/microservice-go-template ${MSROOTTMP}
RUN chmod + ${MSROOTTMP}/sh/prod_build.sh
RUN chmod -R 777 ${MSROOTTMP}

RUN cp ${MSROOTTMP}/sh/prod_build.sh ${MSROOTTMP}
RUN cd "${MSROOTTMP}" && go build -o "${MSROOTTMP}"/build/microservice-go-template "${MSROOTTMP}"/src/main
RUN chmod -R 777 ${MSROOTTMP}/build/

RUN mkdir ${MSROOTOPT}
RUN chmod -R 777 ${MSROOTOPT}
RUN mv ${MSROOTTMP}/build/* ${MSROOTTMP}
RUN mv ${MSROOTTMP}/config/* ${MSROOTTMP}/config
RUN rm -rf ${MSROOTTMP}/*

RUN ${MSROOTTMP}/microservice-go-template
