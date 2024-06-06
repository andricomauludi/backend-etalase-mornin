# syntax=docker/dockerfile:1

FROM golang:1.20

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY *.go ./

# Build
RUN go get github.com/andricomauludi/backend-etalase-mornin/tree/main/controllers/authcontroller
RUN go get github.com/andricomauludi/backend-etalase-mornin/tree/main/controllers/productcontroller
RUN go get github.com/andricomauludi/backend-etalase-mornin/tree/main/controllers/transactioncontroller
RUN go get github.com/andricomauludi/backend-etalase-mornin/tree/main/initializers
RUN go get github.com/andricomauludi/backend-etalase-mornin/tree/main/middleware
RUN go get github.com/andricomauludi/backend-etalase-mornin/tree/main/models

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8090

# Run
CMD ["/docker-gs-ping"]