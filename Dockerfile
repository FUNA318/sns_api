FROM golang:1.24.0

RUN apt update 

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go install github.com/air-verse/air@latest

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
