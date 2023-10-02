FROM golang:1.18 AS build

WORKDIR /app

COPY go.mod ./
COPY . cmd/server/main.go ./

RUN go build -o /server

FROM gcr.io/distroless/base-debian10

ENV TZ="America/Sao_Paulo"

CMD ["/bin/bash"]

WORKDIR /

COPY --from=build /server /server

EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT ["/server"]
