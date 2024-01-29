FROM golang:1.19-alpine
LABEL author="Paul Igweze"

# RUN apk update && apk upgrade && apk add --no-cache bash git
# RUN apk update && \
#     apk upgrade && \
#     apk add --no-cache bash git

ENV SOURCES /go/src/github.com/chibuikeIg/dc/filemetadata-service/

COPY ./file_meta_data_service ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go build

ENV PORT=9091

EXPOSE 9091

WORKDIR ${SOURCES}

CMD ${SOURCES}filemetadata-service