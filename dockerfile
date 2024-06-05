# FROM golang:alpine

# RUN apk update && apk add --no-cache git

# WORKDIR /app

# COPY . .

# RUN go mod tidy

# RUN go build -o binary

# ENTRYPOINT ["/app/binary"]

FROM golang:1.10.5-alpine3.8

 

WORKDIR /go/src/app

COPY main.go .

RUN go build -o main .

EXPOSE 8090

CMD [“./main”]