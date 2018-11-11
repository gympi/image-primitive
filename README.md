# Development
## Run symple project
```
go run *.go
```

# Docker container

## Build container
```
docker build -t primitive_builder .
```

## Run container
```
docker run --publish 9001:9001 --name primitive_builder --rm primitive_builder
```

## Stop container
```
docker stop primitive_builder
```
