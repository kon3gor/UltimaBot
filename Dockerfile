FROM registry.semaphoreci.com/golang:1.19 as builder

ENV APP_HOME /app

WORKDIR "$APP_HOME"
COPY . .

RUN go mod download
RUN go mod verify
RUN go build cmd/ultima/ultima.go

FROM registry.semaphoreci.com/golang:1.19

ENV APP_HOME /app
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

COPY --from=builder "$APP_HOME"/ultima $APP_HOME

CMD ["./ultima"]
