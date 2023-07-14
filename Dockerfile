FROM golang:1.19.3-alpine as dev
RUN mkdir -p /app

COPY . /app
WORKDIR /app
RUN go mod download
RUN go install github.com/cosmtrek/air@v1.29.0
CMD ["air", "-c", ".air.toml"]