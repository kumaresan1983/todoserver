FROM golang:alpine AS build

RUN apk add git

RUN mkdir /src
ADD . /src
WORKDIR /src
COPY . /sbin
RUN go build -o /tmp/todoserver ./cmd/api-server/main.go

FROM alpine:edge

COPY --from=build /tmp/todoserver /sbin/todoserver
COPY --from=build  /sbin/local.env /local.env

CMD /sbin/todoserver
