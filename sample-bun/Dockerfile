FROM golang:1.21

RUN apt-get update && \
    apt-get install -y wget

RUN wget https://github.com/jwilder/dockerize/releases/download/v0.7.0/dockerize-linux-amd64-v0.7.0.tar.gz && \
    tar -C /usr/local/bin -xzvf dockerize-linux-amd64-v0.7.0.tar.gz && \
    rm dockerize-linux-amd64-v0.7.0.tar.gz

RUN wget https://github.com/k0kubun/sqldef/releases/latest/download/mysqldef_linux_amd64.tar.gz && \
    tar -C /usr/local/bin -xzvf mysqldef_linux_amd64.tar.gz && \
    rm mysqldef_linux_amd64.tar.gz

WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]

