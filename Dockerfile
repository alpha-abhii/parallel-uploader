FROM golang:1.25.3-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /parallel-uploader ./cmd/server

EXPOSE 8080

CMD [ "/parallel-uploader" ]