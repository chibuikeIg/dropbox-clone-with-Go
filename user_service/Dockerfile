FROM golang:1.19-alpine
LABEL author="Paul Igweze"

# RUN apk update && apk upgrade && apk add --no-cache bash git
# RUN apk update && \
#     apk upgrade && \
#     apk add --no-cache bash git

ENV SOURCES /go/src/github.com/chibuikeIg/dc/user-service/

COPY ./user_service ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go build

ENV PORT=9090

EXPOSE 9090

WORKDIR ${SOURCES}

CMD ${SOURCES}user-service