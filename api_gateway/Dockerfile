FROM golang:1.19-alpine
LABEL author="Paul Igweze"

# RUN apk update && apk upgrade && apk add --no-cache bash git
# RUN apk update && \
#     apk upgrade && \
#     apk add --no-cache bash git

ENV SOURCES /go/src/github.com/chibuikeIg/dc/api-gateway/

COPY ./api_gateway ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go build

ENV PORT=8080

EXPOSE 8080

WORKDIR ${SOURCES}

CMD ${SOURCES}api-gateway