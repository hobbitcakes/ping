FROM golang:1.23
WORKDIR /app
RUN uname -a
RUN ping google.com -n 4
COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build

EXPOSE 8080

CMD ["./ping"]
