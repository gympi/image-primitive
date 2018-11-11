# Development

## Get or build priject
```shell
export GOPATH=~/go
go get github.com/gympi/image-primitive
```

## Build from source and run
```shell
export GOPATH=~/go

cd $GOPATH/src/

git clone https://github.com/gympi/image-primitive.git

cd ./image-primitive

go install ./...
go run *.go
```

# Docker container

## Build container
```shell
docker build -t primitive_builder .
```

## Run container
```shell
docker run --publish 9001:9001 --name primitive_builder --rm primitive_builder
```

## Stop container
```shell
docker stop primitive_builder
```

# Used third-party projects
## gowebapp
```shell
https://github.com/grisha/gowebapp
```
