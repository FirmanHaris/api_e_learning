FROM golang:1.20.3-bullseye

# RUN apk add git

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/api

EXPOSE 8080/tcp
CMD ["api"]