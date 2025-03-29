# Usar uma imagem base do Go
FROM golang:1.20 AS build

# Definir o diretório de trabalho
WORKDIR /app

# Copiar os arquivos do backend
COPY . .

# Construir o binário do Go
RUN go build -o main .

# Instalar o cliente PostgreSQL
RUN apt-get update && apt-get install -y postgresql-client && rm -rf /var/lib/apt/lists/*

# Copiar o script wait-for-db.sh
COPY wait-for-db.sh /app/wait-for-db.sh
RUN chmod +x /app/wait-for-db.sh

# Configurar o comando de inicialização
ENTRYPOINT ["/app/wait-for-db.sh"]
CMD ["/app/main"]