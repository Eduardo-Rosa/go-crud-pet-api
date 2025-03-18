FROM golang:1.24.0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o go-crud-pet-api .
RUN chmod +x go-crud-pet-api

EXPOSE 8080

CMD ["./go-crud-pet-api"]
