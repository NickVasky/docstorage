FROM golang:1.24.6-alpine3.22 AS go_build

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod tidy
RUN go mod download

COPY . .

RUN go build -o app ./cmd/app.go

FROM alpine:3.22 AS go_run

WORKDIR /usr/src/app

COPY --from=go_build /usr/src/app/app .
COPY --from=go_build /usr/src/app/.env .
