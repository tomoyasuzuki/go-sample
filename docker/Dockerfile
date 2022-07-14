FROM golang:1.18-alpine

ENV ROOT=/go/src/app

WORKDIR ${ROOT}

COPY go.mod ./
RUN go mod download

COPY app ./app
COPY pkg ./pkg

EXPOSE 8080

# Run as rootless container
RUN adduser -D dev && chown -R dev ${ROOT}
USER dev

CMD ["go", "run", "app/main.go"]

# TODO: Use docker-credential-pass

# BUILD: docker build -t go-sample:latest -f docker/Dockerfile . at go-sample/
# RUN  : docker run -it -p 8080:8080 go-sample:latest