FROM golang:1.12.6-alpine

ARG SOURCES="/go/src/github.com/vinny-sabatini/cloud-native-go-intro"

COPY . ${SOURCES}

RUN cd ${SOURCES} && go install

EXPOSE 8080

ENTRYPOINT [ "cloud-native-go-intro" ]