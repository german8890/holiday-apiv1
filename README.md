# holidays-seeker

## Description

This is a simple script to find holidays in a given city.

## Development

### Usage (go ~ docker ~ docker-compose)

```bash
## go executable
$ go run cmd/main.go
```

```bash
## Build the image
$ docker build -t holidays-seeker .

## Run the container with the image
$ docker run -it holidays-seeker
 ```

```bash
## Build the image and run the container (add '-d' for detached mode)
$ docker-compose up --build 
```

### Configs

#### Configuration file format in YAML, loaded from `./internal/shared/config/config.yml` file.

```yaml
server:
  port: 8080

api:
  holiday:
    url: https://farmanet.minsal.cl/index.php/ws/getLocales
```

### Tools

- [Go](https://go.dev/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Fiber](https://gofiber.io/)
- [Viper](https://github.com/spf13/viper)
- [Mock](github.com/golang/mock)
- [Testify](github.com/stretchr/testify)
### Requirements

- Go 1.18
- Docker
- Docker Compose
- Make
- Git

### Endpoints

#### host: `localhost:8080`
#### prefix: `/api/v1`

- `GET /holidays?extra={extra}&type={type}`

### Examples

##### Request for holidays in a given `extra` and type `json`

```bash
$ curl -X GET "http://localhost:8080/api/v1/holidays?extra=Civil&type="
```

#### type=JSON (default)

```json5
{
  "message": "OK",
  "data": [
    {
      "local_nombre": "CRUZ VERDE",
      "comuna_nombre": "Civil",
      "local_direccion": "LOS GINKOS 5 LOCAL 11,12,13",
      "local_telefono": "+56322857355"
    },
    {
      "local_nombre": "CRUZ VERDE",
      "comuna_nombre": "Civil",
      "local_direccion": "AV. CON CON REÑACA 3850 LOCAL 1013",
      "local_telefono": "+56322858104"
    },
  ]
}
````

--

##### Request for holidays in a given `extra` and type `xml`

```bash
$ curl -X GET "http://localhost:8080/api/v1/holidays?extra=Civil&type=xml"
```

#### type=XML

```xml

<Holidays>
    <Holiday>
        <local_nombre>CRUZ VERDE</local_nombre>
        <comuna_nombre>Civil</comuna_nombre>
        <local_direccion>LOS GINKOS 5 LOCAL 11,12,13</local_direccion>
        <local_telefono>+56322857355</local_telefono>
    </Holiday>
    <Holiday>
        <local_nombre>CRUZ VERDE</local_nombre>
        <comuna_nombre>Civil</comuna_nombre>
        <local_direccion>AV. CON CON REÑACA 3850 LOCAL 1013</local_direccion>
        <local_telefono>+56322858104</local_telefono>
    </Holiday>
</Holidays>
```

### Swagger

```shell

# To generate a swagger spec document for a go application
$ swagger generate spec -o ./swagger.json

# Spec validation tool
$ swagger validate https://raw.githubusercontent.com/swagger-api/swagger-spec/master/examples/v2.0/json/petstore-expanded.json

# Generate a client from a swagger spec
$ swagger generate client [-f ./swagger.json] -A [application-name [--principal [principal-name]]

# Generate a server from a swagger spec
$ swagger generate server -f ./swagger.json -A [application-name] [--principal [principal-name]]
```