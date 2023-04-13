FROM golang:1.19.6

COPY . /app
WORKDIR /app

EXPOSE 9090

ENTRYPOINT [ "go", "run", "main.go" ]
