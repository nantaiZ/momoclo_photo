FROM golang:1.15.6-alpine3.12
WORKDIR /go/src/api
EXPOSE 9000 
ENV CGO_ENABLED 0
RUN apk update \
    && apk add --no-cache git \
    && apk add --no-cache mysql-client \
    && go get github.com/oxequa/realize \
    && go get github.com/labstack/echo \
	&& go get github.com/labstack/echo/middleware \
    && go get github.com/jinzhu/gorm \
    && go get github.com/jinzhu/gorm/dialects/mysql
