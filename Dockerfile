FROM golang:1.18-alpine


WORKDIR /app

COPY .  .

RUN go mod tidy

RUN go build -o binary

EXPOSE 8088

CMD [ "/app/binary" ]

RUN apk update && apk add --update make

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest