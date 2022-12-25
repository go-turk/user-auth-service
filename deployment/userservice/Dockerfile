FROM golang:1.19-alpine

WORKDIR /adana

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main /adana/cmd/userservice

CMD ["./main"]