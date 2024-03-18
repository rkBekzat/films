FROM golang:1.21-alpine as builder

WORKDIR /usr/local/src

RUN apk --no-cache add git make

#dependencies
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

#build
COPY . .
RUN go build -o ./bin/app cmd/main.go

FROM alpine as runner

COPY --from=builder /usr/local/src/bin/app /
COPY /configs /configs

CMD ["/app"]