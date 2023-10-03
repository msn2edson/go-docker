FROM golang:1.21.1-alpine3.18 as build

WORKDIR /app

COPY go.mod ./
COPY . cmd/server/main.go ./

RUN go build -o /server

FROM alpine:3.18

ENV TZ="America/Sao_Paulo"

CMD ["/bin/bash"]

WORKDIR /

COPY --from=build /server /server

EXPOSE 3000

ENTRYPOINT ["/server"]
