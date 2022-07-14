# go-sample
Sample repository of API server written in Go
## Build
```
docker build -t go-sample:latest -f docker/Dockerfile .
```
## Run
```
docker run -it -p 8080:8080 go-sample:latest
```

## License
MIT