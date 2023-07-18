FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download && go mod verify

RUN go build -v -o main .
CMD [ "/app/main" ]
