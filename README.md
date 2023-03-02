# holidays-seeker

## Description

This is a script that allows to find the holidays.

## Development

### Usage (go ~ docker ~ docker-compose)

bash
## go executable
$ go run cmd/main.go


bash
## Build the image
$ docker build -t holidays-seeker .

## Run the container with the image
$ docker run -it holidays-seeker
 

bash
## Build the image and run the container (add '-d' for detached mode)
$ docker-compose up --build 


### Configs

#### Configuration file format in YAML, loaded from `./internal/shared/config/config.yml` file.

yaml
server:
  port: 8080

api:
  holiday:
    url: https://api.victorsanmartin.com/feriados/en.json
        


### Tools

- [Go](https://go.dev/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Fiber](https://gofiber.io/)
- [Viper](https://github.com/spf13/viper)
- [Mock](github.com/golang/mock)
- [Testify](github.com/stretchr/testify)
### Requirements

- Go 1.19
- Docker
- Docker Compose
- Make
- Git

### Endpoints

#### host: `localhost:8080`
#### prefix: `/api/v1`

- `GET /holidays?extra={extra}&date={date}`

### Examples

##### Request for holidays in a given `extra` and type `json`

bash
$ curl -X GET "http://localhost:8080/api/v1/holidays/extra?extra=Civil"


#### type=JSON (default)

json5
{
 "status": "success",
  "data": [
    {
      "date": "2023-01-01",
      "title": "Año Nuevo",
      "type": "Civil",
      "inalienable": true,
      "extra": "Civil e Irrenunciable"
    },
    {
      "date": "2023-01-02",
      "title": "Feriado Adicional",
      "type": "Civil",
      "inalienable": false,
      "extra": "Civil"
    },
  ]
}
`

--

##### Request for holidays in a given `extra` and type `xml`

bash
$ curl -X GET "http://localhost:8080/api/v1/holidays/date?date=2023-01-02"
$ curl -X GET  "http://localhost:8080/api/v1/holidays/extra?extra=Civil"


#### type=XML

xml

<Holidays>
    <Holiday>
        <date>2023-01-01</date>
        <title>Año Nuevo</title>
        <type>Civil</type>
        <inalienable>true</inalienable>
        <extra>Civil e Irrenunciable</extra>
    </Holiday>
    <Holiday>
        <date>2023-01-02</date>
        <title>Feriado Adicional</title>
        <type>Civil</type>
        <inalienable>false</inalienable>
        <extra>Civil</extra>
    </Holiday>
</Holidays>


### Swagger

shell

# To generate a swagger spec document for a go application
$ swagger generate spec -o ./swagger.json

# Spec validation tool
$ swagger validate https://raw.githubusercontent.com/swagger-api/swagger-spec/master/examples/v2.0/json/petstore-expanded.json

# Generate a client from a swagger spec
$ swagger generate client [-f ./swagger.json] -A [application-name [--principal [principal-name]]

# Generate a server from a swagger spec
$ swagger generate server -f ./swagger.json -A [application-name] [--principal [principal-name]]
```