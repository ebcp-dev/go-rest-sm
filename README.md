# go-rest-sm
Go (Golang) API REST with Gin Framework

[![Build Status](https://travis-ci.com/ebcp-dev/go-rest-sm.svg?branch=master)](https://travis-ci.com/ebcp-dev/go-rest-sm)


## 1. Run with Docker

1. **Build**

```shell script
make build
docker build . -t api-rest
```

2. **Run**

```shell script
docker run -p 3000:3000 api-rest
```

3. **Test**

```shell script
go test -v ./test/...
```

_______

## 2. Generate Docs

```shell script
# Get swag
go get -u github.com/swaggo/swag/cmd/swag

# Generate docs
export PATH=$(go env GOPATH)/bin:$PATH
swag init --dir cmd/api --parseDependency --output docs
```

Run and go to **http://localhost:3000/docs/index.html**
