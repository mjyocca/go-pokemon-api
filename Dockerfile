# FROM golang:1.18.1 as Builder

# RUN mkdir /app
# ADD . /app
# WORKDIR /app
# RUN CGO_ENABLED=0 GOOS=linux go build -o main ./...

# FROM alpine:latest AS production

# COPY --from=builder /app .

# CMD ["./main"]

FROM golang:1.18.1 as base

FROM base as dev

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /opt/app/api
CMD ["air"]