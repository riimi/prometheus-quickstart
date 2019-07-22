FROM golang:1.12-alpine as build

WORKDIR /go/app
COPY . .
ENV GO111MODULE=on
RUN apk add --no-cache git
RUN go mod download
#RUN CGO_ENABLED=0 go test -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o app

FROM alpine

WORKDIR /app
COPY --from=build /go/app/app .

EXPOSE 1323
CMD ["./app"]