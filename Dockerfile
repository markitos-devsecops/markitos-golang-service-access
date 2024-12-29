FROM golang:1.23 AS build

WORKDIR /app
COPY . /app
COPY ./sample.app.env /go/bin/app.env
RUN go mod download && CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static-debian11:latest-amd64

COPY --from=build /go/bin/app /
COPY --from=build /go/bin/app.env /app.env

ENV SERVICE_PORT 3000
EXPOSE ${SERVICE_PORT}

CMD ["/app"]