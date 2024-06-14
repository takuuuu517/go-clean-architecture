FROM golang:1.22-alpine3.19 as dev

WORKDIR /app
RUN go install github.com/cosmtrek/air@v1.41.0
COPY . /app

RUN apk update && apk add git
RUN git config --global http.postBuffer 5M
RUN go mod tidy || ( go env -w GOPROXY=direct && go mod tidy )

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

EXPOSE 1323

CMD ["air", "-c", ".air.toml"]
