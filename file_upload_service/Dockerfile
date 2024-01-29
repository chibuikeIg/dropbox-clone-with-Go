FROM golang:1.19-alpine
LABEL author="Paul Igweze"

# RUN apk update && apk upgrade && apk add --no-cache bash git
# RUN apk update && \
#     apk upgrade && \
#     apk add --no-cache bash git

ENV SOURCES /go/src/github.com/chibuikeIg/dc/file-upload-service/

COPY ./file_upload_service ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go build

ENV PORT=8081

EXPOSE 8081

WORKDIR ${SOURCES}

CMD ${SOURCES}file-upload-service