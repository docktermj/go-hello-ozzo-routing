# go-hello-ozzo-routing

Part of a comparision of [Go HTTP routers](https://github.com/avelino/awesome-go/blob/master/README.md#routers).

1. https://github.com/go-ozzo
1. https://github.com/go-ozzo/ozzo-routing

## Usage

### Invocation

#### Server

```console
go-hello-ozzo-routing
```

#### Client

In a web browser:

1. Visit http://localhost:8080/ for "Hello world"

`curl` commands

```console
curl -v -X GET http://localhost:8080/streaming
```

## Development

### Dependencies

#### Set environment variables

```console
export GOPATH="${HOME}/go"
export PATH="${PATH}:${GOPATH}/bin:/usr/local/go/bin"
export PROJECT_DIR="${GOPATH}/src/github.com/docktermj"
export REPOSITORY_DIR="${PROJECT_DIR}/go-hello-ozzo-routing"
```

#### Download project

```console
mkdir -p ${PROJECT_DIR}
cd ${PROJECT_DIR}
git clone git@github.com:docktermj/go-hello-ozzo-routing.git
```

#### Download dependencies

```console
cd ${REPOSITORY_DIR}
make dependencies
```

### Build

#### Local build

```console
cd ${REPOSITORY_DIR}
make
```

The results will be in the `${GOPATH}/bin` directory.
