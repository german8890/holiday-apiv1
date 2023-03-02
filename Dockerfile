##
## Build
##
FROM golang:1.19-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /holiday-api cmd/main.go

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /holiday-api /holiday-api

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/holiday-api"]